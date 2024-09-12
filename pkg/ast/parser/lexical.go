package parser

import (
	"errors"
	"vsl/pkg/ast"
)

var LexicalNoSupportError = errors.New("语法不支持")

type LexicalRule interface {
	// Support 是否支持当前语法解析
	Support(parser *TokenParser) bool
	// Parse 语法解析
	Parse(parser *TokenParser) (*ast.Token, error)
}

type LexicalGroup struct {
	rules []LexicalRule
}

func (g *LexicalGroup) Support(parser *TokenParser) bool {
	for i := range g.rules {
		if g.rules[i].Support(parser) {
			return true
		}
	}
	return false
}

func (g *LexicalGroup) Parse(parser *TokenParser) (*ast.Token, error) {
	for i := range g.rules {
		if g.rules[i].Support(parser) {
			return g.rules[i].Parse(parser)
		}
	}
	return nil, LexicalNoSupportError
}

func NewLexicalGroup(rules ...LexicalRule) LexicalRule {
	return &LexicalGroup{rules: rules}
}

type LexicalPrefixRule struct {
	prefix string
	parse  func(parse *TokenParser, rule *LexicalPrefixRule) (*ast.Token, error)
}

func NewLexicalPrefixRule(prefix string, parse func(parse *TokenParser, rule *LexicalPrefixRule) (*ast.Token, error)) *LexicalPrefixRule {
	return &LexicalPrefixRule{prefix: prefix, parse: parse}
}

func (g *LexicalPrefixRule) Support(parser *TokenParser) bool {
	return parser.startWith(0, g.prefix)
}

func (g *LexicalPrefixRule) Parse(parser *TokenParser) (*ast.Token, error) {
	return g.parse(parser, g)
}
