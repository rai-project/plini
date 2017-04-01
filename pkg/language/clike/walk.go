package clike

import (
	"github.com/cznic/cc"
	"github.com/xlab/c-for-go/translator"
)

func Translate(tu *cc.TranslationUnit) (*translator.Translator, error) {

	tl, err := translator.New(nil)
	if err != nil {
		return nil, err
	}
	tl.Learn(tu)

	return tl, nil
}
