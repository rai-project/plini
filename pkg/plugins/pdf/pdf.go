package pdf

import (
	"bytes"
	"errors"
	"io"

	pdfpkg "github.com/puddingfactory/pdf"
)

type pdf struct {
	path      string
	reader    *pdfpkg.Reader
	read      string
	readIndex int64
	done      bool
}

func New(path string) *pdf {
	return &pdf{
		path:   path,
		reader: nil,
		done:   false,
	}
}

func (p *pdf) Open(inputPath string) (*pdfpkg.Reader, error) {
	reader, err := pdfpkg.Open(inputPath)
	if err != nil {
		return nil, err
	}
	p.reader = reader
	return reader, nil
}

const newlineDelimiter = "\n"

func (p *pdf) Read(b []byte) (n int, err error) {
	r := p.reader
	if r == nil {
		return 0, errors.New("pdf file not open")
	}
	if p.done {
		return 0, io.EOF
	}
	p.doRead()
	n = copy(b, p.read[p.readIndex:])
	p.readIndex += int64(n)
	return n, nil
}

func (p *pdf) Close() error {
	return nil
}

func (p *pdf) doRead() error {
	if p.read != "" {
		return nil
	}
	r := p.reader
	var buf bytes.Buffer
	var prev, cur pdfpkg.Text
	for p := 1; p <= r.NumPage(); p++ {
		for _, letter := range r.Page(p).Content().Text {
			prev = cur
			cur = letter

			if prev.S != "" && prev.Y != cur.Y {
				buf.WriteString(newlineDelimiter)
			}
			buf.WriteString(letter.S)
		}
	}
	p.read = buf.String()
	return nil
}
