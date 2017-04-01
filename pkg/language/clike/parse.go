package clike

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cznic/cc"
	"github.com/cznic/ccir"
	"github.com/cznic/mathutil"
	"github.com/cznic/strutil"
	cparser "github.com/xlab/c-for-go/parser"
)

var (
	ccTestdata string
)

func NewConfig(paths ...string) *cparser.Config {
	return &cparser.Config{
		SourcesPaths: paths,
		IncludePaths: []string{"/usr/include"},
	}
}

func Parse1(src []string, opts ...cc.Opt) (string, *cc.TranslationUnit, error) {
	ip, err := cc.ImportPath()
	if err != nil {
		panic(err)
	}

	modelName := fmt.Sprint(mathutil.UintPtrBits)
	model, err := ccir.Model(modelName)
	if err != nil {
		return "", nil, err
	}

	for _, v := range filepath.SplitList(strutil.Gopath()) {
		p := filepath.Join(v, "src", ip, "testdata")
		fi, err := os.Stat(p)
		if err != nil {
			continue
		}

		if fi.IsDir() {
			ccTestdata = p
			break
		}
	}
	if ccTestdata == "" {
		panic("cannot find cc/testdata/")
	}

	ast, err := cc.Parse(
		fmt.Sprintf(`
#define __STDC_HOSTED__ 1
#define __STDC_VERSION__ 199901L
#define __STDC__ 1
`),
		src,
		model,
		opts...,
	)
	return modelName, ast, err
}
