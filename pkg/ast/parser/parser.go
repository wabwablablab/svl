package parser

import (
	"errors"
	"go/token"
	"io"
	"strings"
	"vsl/pkg/ast"
)

// TokenParser Token 解析器
type TokenParser struct {
	// 当前解析到的行号
	line uint
	// 解析到列
	column uint
	// 当前正在处理的数据索引
	index int
	// 当前正在处理得Token索引
	tokenIndex int
	// 所有的Token
	tokens []*token.Token
	// source 源代码
	source string
}

func (p *TokenParser) token() (*ast.Token, error) {
	for {
		code, e := p.code()
		if e != nil {
			return nil, e
		}
		if code == '\n' {
			p.line++
			p.jumpIndex(1)
			p.column = 1
		} else if code == '\r' || code == ' ' || code == '\t' {
			if code == '\r' {
				p.column--
			}
			p.jumpIndex(1)
		} else {
			if globalLexicalRule.Support(p) {
				return globalLexicalRule.Parse(p)
			} else {
				return nil, errors.New("No TOKEN")
			}
		}
	}

}
func (p *TokenParser) startWith(offset int, value string) bool {
	return strings.HasPrefix(p.source[p.index+offset:], value)
}
func (p *TokenParser) jumpIndex(offset int) {
	p.index += offset
	p.column += uint(offset)
}

func (p *TokenParser) text(offsetStart, offsetEnd int) string {
	return p.source[p.index+offsetStart : p.index+offsetEnd]
}

func (p *TokenParser) code() (rune, error) {
	if p.index >= len(p.source) {
		return rune(0), io.EOF
	}
	c := rune(p.source[p.index])
	return c, nil
}
func (p *TokenParser) codeByOffset(offset int) (rune, error) {
	if p.index+offset >= len(p.source) {
		return rune(0), io.EOF
	}
	c := rune(p.source[p.index+offset])
	return c, nil
}

func (p *TokenParser) Read() {

}
