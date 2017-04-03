package example

func init() {
	Dataflow.New().
		ZipWith(
			fs.List("s3://files.p3sr.com/micro/{}.pdf"),
			fs.List("s3://files.p3sr.com/micro/{}_meta.json"),
			fs.FileName,
		).Map(Parallel(
		Step{
			Name:      "cermine",
			Language:  "sh",
			Command:   "java -cp cermine-impl-1.12-jar-with-dependencies.jar pl.edu.icm.cermine.PdfNLMContentExtractor -path {{.Input}}",
			Container: "cli99/cermine",
			Inputs:    []int{0},
		},
		Step{
			Name:      "pdf2txt",
			Language:  "sh",
			Command:   "pdf2txt -q -nopgbrk -enc UTF-8 -eol unix {{.Input}} -",
			Container: "c3sr/pdf2txt",
			Inputs:    []int{0},
		},
		Step{
			Name:      "cermine",
			Language:  "go",
			Command:   "bitbucket.org/c3sr/acm-micro/scrape/Content",
			Container: "c3sr/go",
			Inputs:    []int{0, 1},
		})).
		Join().
		Write("s3://files.p3sr.com/micro/{}_info.json").
		Run()
}
