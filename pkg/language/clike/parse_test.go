package clike

import (
	"io/ioutil"
	"testing"

	"github.com/cznic/cc"
	"github.com/k0kubun/pp"

	cparser "github.com/xlab/c-for-go/parser"
)

const (
	crt0Path = "testdata/include/crt0.c"
)

func TestParseCC(t *testing.T) {
	unit, err := cparser.ParseWith(NewConfig("testdata/hello.c"))
	if err != nil {
		t.Fatal(err)
	}
	pp.Println(unit)

	testUnit(t, unit)
}

func testUnit(t *testing.T, u *cc.TranslationUnit) {
	buf, err := ioutil.ReadFile("testdata/parser_test.out")
	if err != nil {
		t.Fatal(err)
	} else if len(buf) == 0 {
		t.Fatal("no reference output provided")
	}
	if u == nil {
		t.Fatal("no translation unit returned")
	}
	if u.String() != string(buf) {
		log.Println(u)
		t.Error("output doesn't match reference")
	}
}

// func XXTestParseCC(t *testing.T) {
// 	modelName, ast, err := Parse(
// 		[]string{crt0Path},
// 		cc.EnableDefineOmitCommaBeforeDDD(),
// 		cc.ErrLimit(-1),
// 		cc.SysIncludePaths([]string{"testdata/include/"}),
// 	)
// 	if assert.NoError(t, err) {
// 		pp.Println(errStr(err))
// 	}
// 	pp.Println(ast)
// 	pp.Println(modelName)
// }
