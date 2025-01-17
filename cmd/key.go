package cmd

import (
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/openbiox/bget/urlpool"
	vers "github.com/openbiox/bget/versions"
	"github.com/openbiox/butils/archive"
	bexec "github.com/openbiox/butils/exec"
	cio "github.com/openbiox/butils/io"
	log "github.com/openbiox/butils/log"
	cnet "github.com/openbiox/butils/net"
	"github.com/openbiox/butils/slice"
	"github.com/spf13/cobra"
)

var keyVs map[string][]string

var keyCmd = &cobra.Command{
	Use:   "key [key1 key2 key3...]",
	Short: "Can be used to access URLs via a key string.",
	Long:  `Can be used to access URLs via a key string. e.g. 'item' or 'item@version #releaseVersion', : bwa, reffa-defuse@GRCh38 #97. More see here https://github.com/openbiox/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		keyCmdRunOptions(cmd)
	},
}

func parseKeys() (keys []string) {
	if bgetClis.keys != "" && strings.Contains(bgetClis.keys, bgetClis.seperator) {
		keys = strings.Split(bgetClis.keys, bgetClis.seperator)
	} else if bgetClis.keys != "" {
		keys = []string{bgetClis.keys}
	} else if bgetClis.listFile != "" {
		keys = cio.ReadLines(bgetClis.listFile)
	}
	return keys
}

func downloadKey() {
	keys := parseKeys()
	urls, postShellCmd, _ := vers.QueryKeysInfo(keys, &bgetClis.env)
	done := make(map[string][]string)
	var destDirArray []string
	netOpt := setNetParams(&bgetClis)
	sem := make(chan bool, bgetClis.thread)
	for key, v := range urls {
		for i := range v {
			v[i] = preURLFilter(v[i])
			u, _ := url.Parse(v[i])
			v[i] = strings.TrimSpace(u.String())
			if bgetClis.autoPath {
				destDirArray = append(destDirArray, path.Join(bgetClis.downloadDir, key))
			} else {
				destDirArray = append(destDirArray, bgetClis.downloadDir)
			}
		}
		sem <- true
		go func(key string, v []string, destDirArray []string) {
			defer func() {
				<-sem
			}()
			done[key] = cnet.HttpGetURLs(v, destDirArray, netOpt)
		}(key, v, destDirArray)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
	dest := ""
	for key := range done {
		for i := range done[key] {
			args := ""
			dest = done[key][i]
			if len(postShellCmd[key]) > i {
				args = postShellCmd[key][i]
				args = postCmdRender(args, dest)
			}
			if bgetClis.uncompress {
				if err := archive.UnarchiveLog(dest, path.Dir(dest)); err != nil {
					log.Warn(err)
				}
			}
			if args == "" {
				continue
			}
			cmd := exec.Command("sh", "-c", args)
			logPath := ""
			if bgetClis.saveLog {
				logPath = path.Join(bgetClis.logDir, bgetClis.taskID+"_postShellCmd_"+key+".log")
			}
			if args != "" {
				bexec.Shell(cmd, logPath, bgetClis.quiet)
			}
		}
		urlpool.PostKeyCmds(key, done[key], bgetClis.keys)
	}
}

func preURLFilter(url string) string {
	if strings.Contains(url, "doaj.org") {
		req, _ := http.Head(url)
		return req.Request.URL.String()
	}
	return url
}

func postCmdRender(oldCmd string, dest string) (newCmd string) {
	if hasDest, _ := cio.PathExists(dest); !hasDest {
		return ""
	}
	if bgetClis.cmdExtraFromFlag != "" {
		newCmd = oldCmd + " " + bgetClis.cmdExtraFromFlag
	}
	// define your pattern replace
	newCmd = strings.Replace(oldCmd, "{{downloadDir}}", bgetClis.downloadDir, 100)
	newCmd = strings.Replace(newCmd, "{{dest}}", dest, 100)
	newCmd = strings.Replace(newCmd, "{{pdir}}", path.Dir(dest), 100)
	return newCmd
}

func getAllKeys() (keys []string) {
	for i := range urlpool.BgetToolsPool {
		keys = append(keys, urlpool.BgetToolsPool[i].Name)
	}
	for i := range urlpool.BgetFilesPool {
		keys = append(keys, urlpool.BgetFilesPool[i].Name)
	}
	keys = slice.DropSliceDup(keys)
	sort.Strings(keys)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(false)
	table.SetRowSeparator("")
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	tmp := []string{}
	for i := range keys {
		if i%5 == 0 && len(tmp) > 0 {
			table.Append(tmp)
			tmp = []string{}
		} else if i%5 != 0 {
			tmp = append(tmp, keys[i])
		}
	}
	table.Render()
	//fmt.Printf("%s\n", strings.Join(keys, "\n"))
	return keys
}

func keyCmdRunOptions(cmd *cobra.Command) {
	checkArgs(cmd, "key")
	if bgetClis.showVersions && bgetClis.outxt {
		bgetClis.printFormat = "txt"
	} else if bgetClis.showVersions && bgetClis.outjson {
		bgetClis.printFormat = "json"
	} else if bgetClis.showVersions {
		bgetClis.printFormat = "table"
	}
	if bgetClis.printFormat != "" {
		log.Infoln("Featching versions from local and remote website...")
		keys := parseKeys()
		bgetClis.env["printFormat"] = bgetClis.printFormat
		keyVs = vers.QueryKeysVersions(keys, &bgetClis.env)
		return
	}
	if bgetClis.keysAll {
		getAllKeys()
		bgetClis.helpFlags = false
		return
	}
	checkDownloadDir(bgetClis.keys != "")
	if bgetClis.keys != "" || bgetClis.listFile != "" {
		downloadKey()
		bgetClis.helpFlags = false
	}
	if bgetClis.helpFlags {
		cmd.Help()
	}
}

func init() {
	keyCmd.Flags().BoolVar(&(bgetClis.autoPath), "autopath", false, "Logical indicating that whether to create subdir in download dir: e.g. reffa/{{key}}/")
	keyCmd.Flags().BoolVarP(&(bgetClis.showVersions), "show-versions", "v", false, "Show all available versions of key.")
	keyCmd.Flags().BoolVarP(&(bgetClis.outjson), "out-json", "", false, "Output information in JSON string")
	keyCmd.Flags().BoolVarP(&(bgetClis.outxt), "out-text", "", false, "Output information in plain text")
	keyCmd.Flags().BoolVarP(&(bgetClis.keysAll), "keys-all", "a", false, "Show all available string key can be download.")
	keyCmd.Example = `  bget key bwa
  bget key -a // get all available keys
  bget key samtools -v table // view all samtools available versions in CMD table
  bget key samtools -v json // view all samtools available versions in JSON format
  bget key "reffa-defuse@GRCh38 #97" -t 10 -f
  bget key reffa-defuse@GRCh38 release=97 -t 10 -f
  bget key db_annovar@clinvar_20170501 db_annovar@clinvar_20180603 builder=hg38

  bget key db_annovar -v txt
  bget key db_annovar version='clinvar_20131105, clinvar_20140211, clinvar_20140303, clinvar_20140702, clinvar_20140902, clinvar_20140929, clinvar_20150330, clinvar_20150629, clinvar_20151201, clinvar_20160302, clinvar_20161128, clinvar_20170130, clinvar_20170501, clinvar_20170905, clinvar_20180603, avsnp150, avsnp147, avsnp144, avsnp142, avsnp138, cadd, caddgt10, caddgt20, cadd13, cadd13gt10, cadd13gt20, cg69, cg46, cosmic70, cosmic68wgs, cosmic68, cosmic67wgs, cosmic67, cosmic65, cosmic64, dbnsfp35a, dbnsfp33a, dbnsfp31a_interpro, dbnsfp30a, dbscsnv11, eigen, esp6500siv2_ea, esp6500siv2_aa, esp6500siv2_all, exac03nontcga, exac03nonpsych, exac03, fathmm, gerp++gt2, gme, gnomad_exome, gnomad_genome, gwava, hrcr1, icgc21, intervar_20170202, kaviar_20150923, ljb26_all, mcap, mitimpact2, mitimpact24, nci60, popfreq_max_20150413, popfreq_all_20150413, revel, regsnpintron' builder=hg19 -t 10 -f`
}
