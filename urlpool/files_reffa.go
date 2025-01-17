package urlpool

var reffaFiles = []bgetFilesURLType{
	{
		Name: "reffa-ucsc",
		URL:  []string{"https://hgdownload.cse.ucsc.edu/goldenPath/{{version}}/bigZips/{{version}}.fa.gz"},
	},
	{
		Name: "reffa-genecode",
		URL: []string{"http://ftp.ebi.ac.uk/pub/databases/gencode/Gencode_human/release_{{release}}/{{version}}.p12.genome.fa.gz",
			"http://ftp.ebi.ac.uk/pub/databases/gencode/Gencode_human/release_{{release}}/{{version}}_mapping/{{version}}.primary_assembly.genome.fa.gz"},
	},
	{
		Name: "reffa-ensemble",
		URL: []string{"http://ftp.ensembl.org/pub/release-{{release}}/fasta/homo_sapiens/dna/Homo_sapiens.{{version}}.dna.primary_assembly.fa.gz",
			"http://ftp.ensembl.org/pub/release-{{release}}/fasta/homo_sapiens/dna/Homo_sapiens.{{version}}.{{release}}.dna.primary_assembly.fa.gz"},
	},
	{
		Name: "reffa-fusioncatcher",
		URL: []string{"https://svwh.dl.sourceforge.net/project/fusioncatcher/data//human_v{{release}}.tar.gz.aa",
			"https://svwh.dl.sourceforge.net/project/fusioncatcher/data/human_v{{release}}.tar.gz.ab",
			"https://svwh.dl.sourceforge.net/project/fusioncatcher/data/human_v{{release}}.tar.gz.ac",
			"https://svwh.dl.sourceforge.net/project/fusioncatcher/data/human_v{{release}}.tar.gz.ad",
			"https://svwh.dl.sourceforge.net/project/fusioncatcher/data/human_v{{release}}.tar.gz.ae"},
	},
	{
		Name: "reffa-defuse",
		URL: []string{"http://ftp.ensembl.org/pub/release-{{release}}/fasta/homo_sapiens/dna/Homo_sapiens.{{version}}.dna.chromosome.{{chrom}}.fa.gz",
			"http://ftp.ensembl.org/pub/release-{{release}}/gtf/homo_sapiens/Homo_sapiens.{{version}}.{{release}}.gtf.gz",
			"http://hgdownload.cse.ucsc.edu/goldenPath/{{version}}/database/rmsk.txt.gz",
			"http://hgdownload.cse.ucsc.edu/goldenPath/{{version}}/bigZips/est.fa.gz",
			"http://hgdownload.cse.ucsc.edu/goldenPath/{{version}}/database/intronEst.txt.gz",
			"http://ftp.ncbi.nlm.nih.gov/repository/UniGene/Homo_sapiens/Hs.seq.uniq.gz",
			"http://hgdownload.cse.ucsc.edu/goldenPath/{{version}}/database/rmsk.txt.gz"},
	},
	{
		Name: "reffa-encode-hg19",
		URL:  []string{"http://hgdownload.cse.ucsc.edu/goldenpath/hg19/encodeDCC/referenceSequences/male.hg19.fa.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/hg19/hg19.chrom.sizes", "https://storage.googleapis.com/encode-pipeline-genome-data/hg19/wgEncodeDacMapabilityConsensusExcludable.bed.gz"},
	},
	{
		Name: "reffa-encode-hg19-bowtie2-index",
		URL:  []string{"https://storage.googleapis.com/encode-pipeline-genome-data/hg19/bowtie2_index/male.hg19.fa.tar"},
	},
	{
		Name: "reffa-encode-hg19-bwa-index",
		URL:  []string{"https://storage.googleapis.com/encode-pipeline-genome-data/hg19/bwa_index/male.hg19.fa.tar"},
	},
	{
		Name: "reffa-encode-hg19-ataqc",
		URL: []string{
			"https://storage.googleapis.com/encode-pipeline-genome-data/hg19/ataqc/hg19_gencode_tss_unique.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/hg19/ataqc/reg2map_honeybadger2_dnase_all_p10_ucsc.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/hg19/ataqc/reg2map_honeybadger2_dnase_prom_p2.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/hg19/ataqc/reg2map_honeybadger2_dnase_enh_p2.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/hg19/ataqc/dnase_avgs_reg2map_p10_merged_named.pvals.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/hg19/ataqc/eid_to_mnemonic.txt"},
	},
	{
		Name: "reffa-encode-hg38",
		URL:  []string{"https://www.encodeproject.org/files/GRCh38_no_alt_analysis_set_GCA_000001405.15/@@download/GRCh38_no_alt_analysis_set_GCA_000001405.15.fasta.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/hg38/hg38.chrom.sizes", "https://storage.googleapis.com/encode-pipeline-genome-data/hg38/hg38.blacklist.bed.gz"},
	},
	{
		Name: "reffa-encode-hg38-bowtie2-index",
		URL:  []string{"https://storage.googleapis.com/encode-pipeline-genome-data/hg38/bowtie2_index/GRCh38_no_alt_analysis_set_GCA_000001405.15.fasta.tar"},
	},
	{
		Name: "reffa-encode-hg38-bwa-index",
		URL:  []string{"https://storage.googleapis.com/encode-pipeline-genome-data/hg38/bwa_index/GRCh38_no_alt_analysis_set_GCA_000001405.15.fasta.tar"},
	},
	{
		Name: "reffa-encode-hg38-ataqc",
		URL:  []string{"https://storage.googleapis.com/encode-pipeline-genome-data/hg38/ataqc/hg38_gencode_tss_unique.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/hg38/ataqc/reg2map_honeybadger2_dnase_all_p10_ucsc.hg19_to_hg38.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/hg38/ataqc/reg2map_honeybadger2_dnase_prom_p2.hg19_to_hg38.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/hg38/ataqc/reg2map_honeybadger2_dnase_enh_p2.hg19_to_hg38.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/hg38/ataqc/hg38_celltype_compare_subsample.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/hg38/ataqc/hg38_dnase_avg_fseq_signal_formatted.txt.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/hg38/ataqc/hg38_dnase_avg_fseq_signal_metadata.txt"},
	},
	{
		Name: "reffa-encode-mm10",
		URL:  []string{"https://www.encodeproject.org/files/mm10_no_alt_analysis_set_ENCODE/@@download/mm10_no_alt_analysis_set_ENCODE.fasta.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/mm10/mm10.chrom.sizes", "https://storage.googleapis.com/encode-pipeline-genome-data/mm10/mm10.blacklist.bed.gz"},
	},
	{
		Name: "reffa-encode-mm10-bowtie2-index",
		URL:  []string{"https://storage.googleapis.com/encode-pipeline-genome-data/mm10/bowtie2_index/mm10_no_alt_analysis_set_ENCODE.fasta.tar"},
	},
	{
		Name: "reffa-encode-mm10-bwa-index",
		URL:  []string{"https://storage.googleapis.com/encode-pipeline-genome-data/mm10/bwa_index/mm10_no_alt_analysis_set_ENCODE.fasta.tar"},
	},
	{
		Name: "reffa-encode-mm10-ataqc",
		URL:  []string{"https://storage.googleapis.com/encode-pipeline-genome-data/mm10/ataqc/mm10_gencode_tss_unique.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/mm10/ataqc/mm10_univ_dhs_ucsc.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/mm10/ataqc/tss_mm10_master.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/mm10/ataqc/mm10_enh_dhs_ucsc.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/mm10/ataqc/mm10_celltype_compare_subsample.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/mm10/ataqc/mm10_dnase_avg_fseq_signal_formatted.txt.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/mm10/ataqc/mm10_dnase_avg_fseq_signal_metadata.txt"},
	},
	{
		Name: "reffa-encode-mm9",
		URL:  []string{"https://storage.googleapis.com/encode-pipeline-genome-data/mm9/mm9.fa.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/mm9/mm9.chrom.sizes", "https://storage.googleapis.com/encode-pipeline-genome-data/mm9/mm9-blacklist.bed.gz"},
	},
	{
		Name: "reffa-encode-mm9-bowtie2-index",
		URL:  []string{"https://storage.googleapis.com/encode-pipeline-genome-data/mm9/bowtie2_index/mm9.fa.tar"},
	},
	{
		Name: "reffa-encode-mm9-bwa-index",
		URL:  []string{"https://storage.googleapis.com/encode-pipeline-genome-data/mm9/bwa_index/mm9.fa.tar"},
	},
	{
		Name: "reffa-encode-mm9-ataqc",
		URL:  []string{"https://storage.googleapis.com/encode-pipeline-genome-data/mm9/ataqc/mm9_gencode_tss_unique.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/mm9/ataqc/mm9_univ_dhs_ucsc.from_mm10.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/mm9/ataqc/tss_mm9_master.from_mm10.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/mm9/ataqc/mm9_enh_dhs_ucsc.from_mm10.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/mm9/ataqc/mm9_dhs_universal_ucsc_v1.bed.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/mm9/ataqc/dnase_avgs_merged_named.fseq.vals.gz", "https://storage.googleapis.com/encode-pipeline-genome-data/mm9/ataqc/accession_to_name.txt"},
	},
}
