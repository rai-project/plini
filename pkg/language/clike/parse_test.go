package clike

import (
	"fmt"
	"testing"

	rice "github.com/GeertJohan/go.rice"
	"github.com/stretchr/testify/assert"
)

var (
	box   = rice.MustFindBox("testdata")
	test2 = box.MustString("test2.c")
)

func TestParse(t *testing.T) {
	ast, err := Parse(test2)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	fmt.Println(ast)
}
