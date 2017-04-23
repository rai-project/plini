package unidoc

import (
	"os"

	unipdf "github.com/unidoc/unidoc/pdf"
)

type pdf struct {
	path      string
	pdfReader *unipdf.PdfReader
}

func New(path string) *pdf {
	return &pdf{
		path: path,
	}
}

func (p *pdf) Open(inputPath string) (*unipdf.PdfReader, error) {

	f, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	pdfReader, err := unipdf.NewPdfReader(f)
	if err != nil {
		return nil, err
	}

	p.pdfReader = pdfReader

	return pdfReader, nil
}

// func (p *pdf) Read() {
// 	p.pdfReader.GetPage(1).String()
// }

func (p *pdf) Close() error {
	return nil
}
