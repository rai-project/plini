package clike

import (
	"fmt"
	"testing"

	"os"

	"github.com/cznic/cc"
	"github.com/cznic/ccir"
	"github.com/cznic/ir"
	"github.com/k0kubun/pp"
	"github.com/stretchr/testify/assert"
	"github.com/xlab/c-for-go/translator"
)

const (
	crt0Path = "testdata/include/crt0.c"
)

func XXXTestParseCC(t *testing.T) {
	_, tu, err := Parse("testdata/hello.c")
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	// pp.Println(tu)
	_ = tu

	// testUnit(t, tu)
}

func XXTestParse(t *testing.T) {
	modelName, unit, err := Parse("testdata/test2.c")
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	// for ii, _ := range tu.Declarations.Identifiers {
	// 	p := tu.Declarations.Lookup(cc.NSIdentifiers, ii)
	// 	p.Node
	// 	fmt.Println(cc.PrettyString(p))
	// }
	// cc.PrettyString(unit.TranslationUnit)
	// mn, s := unit.Declarations.Lookup2(cc.NSIdentifiers, xc.Dict.SID("main"))
	// e, ok := mn.Node.(*cc.DirectDeclarator)
	// if !assert.True(t, ok) {
	// 	t.Fatal("not ok")
	// }
	// d := cc.PrettyString(e.TopDeclarator())
	// switch g := s.Scope() {
	// case cc.ScopeFile:
	// pp.Println(g)
	// case cc.ScopeBlock:
	// 	if d.Type.Specifier().IsStatic() {
	// 		pp.Println(g)
	// 	}
	// }

	objs, err := ccir.New(modelName, unit)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	for i, v := range objs {
		switch x := v.(type) {
		case *ir.DataDefinition:
			fmt.Fprintf(os.Stdout, "# [%v]: %T %v %v\n", i, x, x.ObjectBase, x.Value)
		case *ir.FunctionDefinition:
			fmt.Fprintf(os.Stdout, "# [%v]: %T %v %v\n", i, x, x.ObjectBase, x.Arguments)
			for i, v := range x.Body {
				fmt.Fprintf(os.Stdout, "%#05x\t%v\n", i, v)
			}
		default:
			t.Fatalf("[%v] %T %v", i, x, x)
		}
	}
}

func XXXXXTestParse(t *testing.T) {
	_, unit, err := Parse("testdata/test2.c")
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	tl, err := Translate(unit)
	if !assert.NoError(t, err) {
		return
	}

	fs, ok := tl.Declares()[0].Spec.(*translator.CFunctionSpec)
	assert.True(t, ok)
	pp.Println(fs)

	pp.Println(tl.ValueMap())

	pp.Println(tl.Declares()[0].Spec)

	// g, ok := decls[0].Value.(*translator.FunctionDefinition)
	// pp.Println(g)

}

// func testUnit(t *testing.T, u *cc.TranslationUnit) {
// 	buf, err := ioutil.ReadFile("testdata/parser_test.out")
// 	if err != nil {
// 		t.Fatal(err)
// 	} else if len(buf) == 0 {
// 		t.Fatal("no reference output provided")
// 	}
// 	if u == nil {
// 		t.Fatal("no translation unit returned")
// 	}
// 	if u.String() != string(buf) {
// 		log.Println(u)
// 		t.Error("output doesn't match reference")
// 	}
// }

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

var tagMap map[cc.Specifier]interface{}

func TestParse(t *testing.T) {
	_, unit, err := Parse("testdata/test2.c")
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	// for ii, _ := range unit.Declarations.Identifiers {
	// 	p := unit.Declarations.Lookup(cc.NSIdentifiers, ii)
	// 	g, _ := p.Node.(*cc.DirectDeclarator)
	// 	fmt.Println(cc.PrettyString(g))

	// }
	ext := unit.TranslationUnit.ExternalDeclaration.FunctionDefinition.FunctionBody.
		CompoundStatement.BlockItemListOpt.BlockItemList.BlockItemList.BlockItem.
		Statement.ExpressionStatement.ExpressionListOpt.ExpressionList.Expression
	fmt.Println(cc.PrettyString(ext))

	// switch ext.Case {
	// case 0: /// FunctionDefinition
	// 	if declr := ext.FunctionDefinition.Declarator; declr != nil {
	// 		// tagMap[declr.RawSpecifier()] = declr
	// 		fmt.Println(cc.PrettyString(declr.DirectDeclarator))
	// 	} else {
	// 		return
	// 	}
	// }
}
