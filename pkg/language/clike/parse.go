package clike

import (
	"fmt"

	"github.com/cznic/cc"
	"github.com/cznic/mathutil"
	cparser "github.com/xlab/c-for-go/parser"
)

var (
	ccTestdata string
)

func newConfig(paths ...string) *cparser.Config {
	return &cparser.Config{
		SourcesPaths: paths,
		CCDefs:       false,
		CCIncl:       false,
	}
}

func Parse1(files ...string) (string, *cc.TranslationUnit, error) {
	conf := newConfig(files...)
	a, b := cparser.ParseWith(conf)
	return "", a, b
}

func Parse(files ...string) (string, *cc.TranslationUnit, error) {
	modelName := fmt.Sprint(mathutil.UintPtrBits)
	predefined := builtinBase
	predefined += basePredefines
	if archDefs, ok := archPredefines[archBits]; ok {
		predefined += fmt.Sprintf("\n%s", archDefs)
	}

	predefined = ""

	ast, err := cc.Parse(
		predefined,
		files,
		model,
		cc.AllowCompatibleTypedefRedefinitions(),
		cc.EnableAlignOf(),
		cc.EnableAlternateKeywords(),
		cc.EnableAnonymousStructFields(),
		cc.EnableAsm(),
		cc.EnableBuiltinClassifyType(),
		cc.EnableBuiltinConstantP(),
		cc.EnableComputedGotos(),
		cc.EnableDefineOmitCommaBeforeDDD(),
		cc.EnableEmptyDeclarations(),
		cc.EnableEmptyStructs(),
		cc.EnableImaginarySuffix(),
		cc.EnableImplicitFuncDef(),
		cc.EnableImplicitIntType(),
		cc.EnableLegacyDesignators(),
		cc.EnableOmitConditionalOperand(),
		cc.EnableOmitFuncArgTypes(),
		cc.EnableOmitFuncRetType(),
		cc.EnableParenthesizedCompoundStatemen(),
		cc.EnableTypeOf(),
		cc.EnableUnsignedEnums(),
		cc.EnableWideBitFieldTypes(),
		cc.ErrLimit(-1),
		cc.SysIncludePaths([]string{"testdata/include/"}),
	)
	return modelName, ast, err
}
