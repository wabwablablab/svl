package parser

import "vsl/pkg/ast"

var globalLexicalRule LexicalRule

func init() {
	globalLexicalRule = NewLexicalGroup(
		NewLexicalGroup(
			NewLexicalPrefixRule("++", simpleParse(ast.TokenKindPlusPlus)),
			NewLexicalPrefixRule("+=", simpleParse(ast.TokenKindPlusAssign)),
			NewLexicalPrefixRule("+", simpleParse(ast.TokenKindPlus)),
			NewLexicalPrefixRule("--", simpleParse(ast.TokenKindSubSub)),
			NewLexicalPrefixRule("-=", simpleParse(ast.TokenKindSubAssign)),
			NewLexicalPrefixRule("-", simpleParse(ast.TokenKindSub)),
			NewLexicalPrefixRule("*=", simpleParse(ast.TokenKindMultiAssign)),
			NewLexicalPrefixRule("*", simpleParse(ast.TokenKindMulti)),
			NewLexicalPrefixRule("/=", simpleParse(ast.TokenKindDivAssign)),
			NewLexicalPrefixRule("/", simpleParse(ast.TokenKindDiv)),
			NewLexicalPrefixRule("%=", simpleParse(ast.TokenKindRemAssign)),
			NewLexicalPrefixRule("%", simpleParse(ast.TokenKindRem)),
			NewLexicalPrefixRule("||", simpleParse(ast.TokenKindOr)),
			NewLexicalPrefixRule("|=", simpleParse(ast.TokenKindBitOrAssign)),
			NewLexicalPrefixRule("|", simpleParse(ast.TokenKindBitOr)),
			NewLexicalPrefixRule("&&", simpleParse(ast.TokenKindAnd)),
			NewLexicalPrefixRule("&=", simpleParse(ast.TokenKindBitAndAssign)),
			NewLexicalPrefixRule("&", simpleParse(ast.TokenKindBitAnd)),
			NewLexicalPrefixRule(">>=", simpleParse(ast.TokenKindBitMoveRightAssign)),
			NewLexicalPrefixRule(">>", simpleParse(ast.TokenKindBitMoveRight)),
			NewLexicalPrefixRule("<<=", simpleParse(ast.TokenKindBitMoveLeftAssign)),
			NewLexicalPrefixRule("<<", simpleParse(ast.TokenKindBitMoveLeft)),
			NewLexicalPrefixRule("==", simpleParse(ast.TokenKindEquals)),
			NewLexicalPrefixRule("=", simpleParse(ast.TokenKindAssign)),
			NewLexicalPrefixRule(">=", simpleParse(ast.TokenKindGreaterEquals)),
			NewLexicalPrefixRule(">", simpleParse(ast.TokenKindGreater)),
			NewLexicalPrefixRule("<=", simpleParse(ast.TokenKindLessEquals)),
			NewLexicalPrefixRule("<", simpleParse(ast.TokenKindLess)),
			NewLexicalPrefixRule("!=", simpleParse(ast.TokenKindNoEquals)),
			NewLexicalPrefixRule("!", simpleParse(ast.TokenKindExclamation)),
		), // 运算符
		NewLexicalGroup(
			NewLexicalPrefixRule("#", noteParse), // 解析注释
			NewLexicalPrefixRule("(", simpleParse(ast.TokenKindBracketStart)),
			NewLexicalPrefixRule(")", simpleParse(ast.TokenKindBracketEnd)),
			NewLexicalPrefixRule(".", simpleParse(ast.TokenKindDot)),
			NewLexicalPrefixRule(",", simpleParse(ast.TokenKindComma)),
			NewLexicalPrefixRule("[", simpleParse(ast.TokenKindSquareBracketStart)),
			NewLexicalPrefixRule("]", simpleParse(ast.TokenKindSquareBracketEnd)),
			NewLexicalPrefixRule("{", simpleParse(ast.TokenKindCurlyBracketStart)),
			NewLexicalPrefixRule("}", simpleParse(ast.TokenKindCurlyBracketEnd)),
			NewLexicalPrefixRule(";", simpleParse(ast.TokenKindBranch)),
		), // 分割符
		&numberParse{}, // 解析数字的
		&stringParse{},
		&identParse{},
	)
}
