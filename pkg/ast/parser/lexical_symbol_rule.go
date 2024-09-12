package parser

import (
	"errors"
	"vsl/pkg/ast"
)

type tokenError struct {
	message string
	line    uint
	column  uint
}

func (t *tokenError) Error() string {
	return t.message
}

func (t *tokenError) RuntimeError() {
	panic(t.message)
}

var TokenNumberError = errors.New("解析数字Token失败")
var TokenStringError = errors.New("解析字符串失败")

// simpleParse
// 简单的解析，解析内容就是前缀内容
func simpleParse(kind ast.TokenKind) func(parse *TokenParser, rule *LexicalPrefixRule) (*ast.Token, error) {
	return func(parse *TokenParser, rule *LexicalPrefixRule) (*ast.Token, error) {
		tok := &ast.Token{
			Value: rule.prefix,
			Line:  parse.line,
			Col:   parse.column,
			Kind:  kind,
		}
		parse.jumpIndex(len(rule.prefix))
		return tok, nil
	}
}

func noteParse(parse *TokenParser, rule *LexicalPrefixRule) (*ast.Token, error) {
	count := 0
	for {
		c, err := parse.codeByOffset(count)
		if err != nil {
			return nil, err
		}
		if c == '\n' {
			break
		}
		count++
	}
	tok := &ast.Token{
		Value: parse.text(0, count),
		Kind:  ast.TokenKindNotes,
		Col:   parse.column,
		Line:  parse.line,
	}
	parse.jumpIndex(count)
	return tok, nil
}

type numberParse struct {
}

func (n *numberParse) Support(parser *TokenParser) bool {
	c, e := parser.code()
	if e != nil {
		return false
	}
	if !(c >= '0' && c <= '9') {
		return false
	}
	return true
}

func (n *numberParse) Parse(parser *TokenParser) (*ast.Token, error) {
	index := 0
	var err error
	var c rune
	for {
		c, err = parser.codeByOffset(index)
		if err != nil {
			return nil, err
		}
		if c >= '0' && c <= '9' {
			index++
		} else {
			break
		}
	}
	value := parser.text(0, index)
	st := &ast.Token{
		Value: value,
		Kind:  ast.TokenKindNumber,
		Line:  parser.line,
		Col:   parser.column,
	}
	parser.jumpIndex(index)
	return st, nil
}

type stringParse struct {
}

func (i *stringParse) Support(parser *TokenParser) bool {
	c, e := parser.code()
	if e != nil {
		return false
	}
	if c == '"' {
		return true
	}
	return false
}

type StringMode = int

const (
	StringStart StringMode = 1 + iota
	StringContent
	StringEnd
)

func (i *stringParse) Parse(parser *TokenParser) (*ast.Token, error) {
	index := 0
	var err error
	var c rune
	escape := false
	model := StringStart
	for {
		c, err = parser.codeByOffset(index)
		if err != nil {
			return nil, err
		}
		if model == StringStart && c == '"' {
			model = StringContent
		} else if !escape && c == '\\' {
			escape = true
		} else if model == StringContent && !escape && c == '"' {
			model = StringEnd
			index++
			// 解析字符串完成
			break
		} else if !escape && c == '\n' {
			// 有问题
			return nil, &tokenError{
				message: "解析字符串失败，缺少关闭\"符号",
				line:    parser.line,
				column:  parser.column + uint(index),
			}
		} else {
			escape = false
		}
		index++
	}
	value := parser.text(0, index)
	st := &ast.Token{
		Value: value,
		Kind:  ast.TokenKindString,
		Line:  parser.line,
		Col:   parser.column,
	}
	parser.jumpIndex(index)
	return st, nil
}

// 标识符不能是数字开头
type identParse struct {
}

func (i *identParse) Support(parser *TokenParser) bool {
	c, e := parser.code()
	if e != nil {
		return false
	}
	if c != ' ' && c != '\n' {
		return true
	}
	return false
}

func (i *identParse) Parse(parser *TokenParser) (token *ast.Token, err error) {
	index := 0
	var c rune
	for {
		c, err = parser.codeByOffset(index)
		if err != nil {
			return nil, err
		}
		if c == ' ' || c == '\n' || c == '\t' || c == '\r' {
			break
		}
		if c == ',' || c == '.' || c == ';' {
			break
		}
		if c == '=' || c == '+' || c == '-' || c == '*' || c == '/' || c == '%' {
			break
		}
		if c == '>' || c == '<' || c == '!' || c == '|' || c == '&' || c == '^' {
			break
		}
		if c == '`' || c == '#' || c == '"' || c == '\'' {
			break
		}
		if c == '(' || c == ')' || c == '{' || c == '}' || c == '[' || c == ']' {
			break
		}
		index++
	}
	value := parser.text(0, index)
	if index <= 0 {
		return nil, &tokenError{
			message: "标识符解析错误：为空",
			line:    parser.line,
			column:  parser.column,
		}
	}

	kind := ast.TokenKindIdent
	if v, ok := ast.TokenKindKeyword[value]; ok {
		kind = v
	}
	st := &ast.Token{
		Value: value,
		Kind:  kind,
		Line:  parser.line,
		Col:   parser.column,
	}
	parser.jumpIndex(index)
	return st, nil
}
