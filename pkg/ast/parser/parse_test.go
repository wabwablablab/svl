package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"vsl/pkg/ast"
)

func TestString(t *testing.T) {
	bytes, err := os.ReadFile("E:\\YunkeAnXin\\VSL\\pkg\\ast\\parser\\data.svl")
	if err != nil {
		return
	}
	content := string(bytes)
	tp := &TokenParser{source: content, line: 1, column: 1}
	for {
		token, err := tp.token()
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		marshal, err := json.Marshal(token)
		if err != nil {
			return
		}
		fmt.Printf("%s => %v\n", ast.TokenKindToString(token.Kind), string(marshal))
	}
}
