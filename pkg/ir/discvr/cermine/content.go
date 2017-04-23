package cermine

import (
	"encoding/xml"
	"os/exec"
	"path/filepath"

	sourcepath "github.com/GeertJohan/go-sourcepath"
	"github.com/apex/log"
)

const (
	jarBase = "cermine-impl-1.12-jar-with-dependencies.jar"
)

var (
	jar string
)

type Content struct {
}

func hasJava() error {
	_, err := exec.LookPath("java")
	return err
}

func run(pdfPath string) (*Article, error) {
	if err := hasJava(); err != nil {
		log.WithError(err).Error("failed in has java")
		return nil, err
	}
	cmd := exec.Command("java", "-cp", jar, "pl.edu.icm.cermine.PdfNLMContentExtractor", "-path", pdfPath)
	buf, err := cmd.Output()
	if err != nil {
		log.WithError(err).Error("failed to process the pdf")
		return nil, err
	}

	res := &Article{}

	if err := xml.Unmarshal(buf, res); err != nil {
		log.WithError(err).Error("Cannot unmarshal cermine output")
		return nil, err
	}
	return res, nil
}

func init() {
	jar = filepath.Join(sourcepath.MustAbsoluteDir(), jarBase)
}
