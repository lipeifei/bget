package cmd

import (
	"path"
	"strings"

	"github.com/openbiox/butils/archive"
	cio "github.com/openbiox/butils/io"
	log "github.com/openbiox/butils/log"
	cnet "github.com/openbiox/butils/net"
	"github.com/spf13/cobra"
)

var urlCmd = &cobra.Command{
	Use:   "url [url1 url2 url3...]",
	Short: "Can be used to access URLs via Golang http, wget, curl, axel and git, and rsync.",
	Long:  `Can be used to access URLs via Golang http, wget, curl, axel and git, and rsync. More see here https://github.com/openbiox/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		urlCmdRunOptions(cmd)
	},
}

func downloadUrls() {
	urls := []string{}
	if bgetClis.urls != "" && strings.Contains(bgetClis.urls, bgetClis.seperator) {
		urls = strings.Split(bgetClis.urls, bgetClis.seperator)
	} else if bgetClis.urls != "" {
		urls = []string{bgetClis.urls}
	} else if bgetClis.listFile != "" {
		urls = cio.ReadLines(bgetClis.listFile)
	}
	var destDirArray []string
	for i := range urls {
		urls[i] = strings.TrimSpace(urls[i])
		if !strings.Contains(urls[i], "://") && !strings.Contains(urls[i], "git@") {
			urls[i] = "https://" + urls[i]
		}
		destDirArray = append(destDirArray, bgetClis.downloadDir)
	}

	netOpt := setNetParams(&bgetClis)
	done := cnet.HttpGetURLs(urls, destDirArray, netOpt)
	for _, dest := range done {
		if bgetClis.uncompress {
			if err := archive.UnarchiveLog(dest, path.Dir(dest)); err != nil {
				log.Warn(err)
			}
		}
	}
}
func urlCmdRunOptions(cmd *cobra.Command) {
	checkArgs(cmd, "url")
	checkDownloadDir(bgetClis.urls != "" || bgetClis.listFile != "")
	if bgetClis.urls != "" || bgetClis.listFile != "" {
		downloadUrls()
		bgetClis.helpFlags = false
	}
	if bgetClis.helpFlags {
		cmd.Help()
	}
}
func init() {
	urlCmd.Example = `  urls="https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe,http://download.oray.com/pgy/windows/PgyVPN_4.1.0.21693.exe,https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe" && echo $urls | tr "," "\n"> /tmp/urls.list

  bget url ${urls}
  bget url https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe
  bget url ${urls} -t 2 -o /tmp/download
  bget url ${urls} -t 3 -o /tmp/download -f -g wget
  bget url ${urls} -t 3 -o /tmp/download -g wget --ignore
  bget url -l /tmp/urls.list -o /tmp/download -f -t 3`
}
