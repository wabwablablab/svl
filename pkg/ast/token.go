package ast

type TokenKind = uint

const (
	begin TokenKind = iota
	// TokenKindNumber 数字，包含 1,1.1
	TokenKindNumber
	// TokenKindString 字符串 以"开头，以"结尾，或者以 `开头 以 `结尾
	TokenKindString
	// TokenKindIdent 标识符 不是数字和字符串得字母
	TokenKindIdent
	// TokenKindNotes 注释开头 #
	TokenKindNotes
	// TokenKindBracketStart 括号开始
	TokenKindBracketStart
	// TokenKindBracketEnd 括号结束
	TokenKindBracketEnd
	// TokenKindDot .
	TokenKindDot
	// TokenKindComma ,
	TokenKindComma
	// TokenKindSquareBracketStart [
	TokenKindSquareBracketStart
	// TokenKindSquareBracketEnd ]
	TokenKindSquareBracketEnd
	// TokenKindCurlyBracketStart {
	TokenKindCurlyBracketStart
	// TokenKindCurlyBracketEnd }
	TokenKindCurlyBracketEnd

	// TokenKindBranch 分号 ;
	TokenKindBranch

	// TokenKindAssign 赋值符号 =
	TokenKindAssign
	// TokenKindExclamation 感叹号
	TokenKindExclamation

	// TokenKindPlus +
	TokenKindPlus
	// TokenKindPlusAssign +=
	TokenKindPlusAssign
	// TokenKindPlusPlus ++
	TokenKindPlusPlus
	// TokenKindSub -
	TokenKindSub
	// TokenKindSubAssign -=
	TokenKindSubAssign
	// TokenKindSubSub --
	TokenKindSubSub
	// TokenKindMulti *
	TokenKindMulti
	// TokenKindMultiAssign *=
	TokenKindMultiAssign
	// TokenKindDiv /
	TokenKindDiv
	// TokenKindDivAssign /=
	TokenKindDivAssign
	// TokenKindRem %
	TokenKindRem
	// TokenKindRemAssign %=
	TokenKindRemAssign
	// TokenKindBitMoveLeft 位运算 <<
	TokenKindBitMoveLeft
	// TokenKindBitMoveLeftAssign <<=
	TokenKindBitMoveLeftAssign
	// TokenKindBitMoveRight 位运算 >>
	TokenKindBitMoveRight
	// TokenKindBitMoveRightAssign >>=
	TokenKindBitMoveRightAssign
	// TokenKindBitOr |
	TokenKindBitOr
	// TokenKindBitOrAssign |=
	TokenKindBitOrAssign
	// TokenKindBitAnd &
	TokenKindBitAnd
	// TokenKindBitAndAssign &=
	TokenKindBitAndAssign

	// TokenKindGreater 大于号
	TokenKindGreater
	// TokenKindGreaterEquals 大等于号
	TokenKindGreaterEquals
	// TokenKindLess 小于号
	TokenKindLess
	// TokenKindLessEquals 小等于号
	TokenKindLessEquals
	// TokenKindEquals 等于符号
	TokenKindEquals
	// TokenKindNoEquals !=
	TokenKindNoEquals

	// TokenKindAnd &&
	TokenKindAnd
	// TokenKindOr ||
	TokenKindOr

	TokenKindIf
	TokenKindElse
	TokenKindFor
	TokenKindContinue
	TokenKindBreak
	TokenKindFunc
	TokenKindReturn
	TokenKindSwitch
	TokenKindCase
	TokenKindDefault
	TokenKindImport
	TokenKindExport
	TokenKindInclude

	TokenKindBase64
	TokenKindHex
	end
)

var tokenKindLabel = map[TokenKind]string{
	TokenKindNumber:             "number",
	TokenKindString:             "string",
	TokenKindIdent:              "TokenKindIdent",
	TokenKindBracketStart:       "TokenKindBracketStart",
	TokenKindBracketEnd:         "TokenKindBracketEnd",
	TokenKindDot:                "TokenKindDot",
	TokenKindComma:              "TokenKindComma",
	TokenKindSquareBracketStart: "TokenKindSquareBracketStart",
	TokenKindSquareBracketEnd:   "TokenKindSquareBracketEnd",
	TokenKindCurlyBracketStart:  "TokenKindCurlyBracketStart",
	TokenKindCurlyBracketEnd:    "TokenKindCurlyBracketEnd",
	TokenKindBranch:             "TokenKindBranch",
	TokenKindAssign:             "TokenKindAssign",
	TokenKindExclamation:        "TokenKindExclamation",
	TokenKindPlus:               "TokenKindPlus",
	TokenKindPlusAssign:         "TokenKindPlusAssign",
	TokenKindPlusPlus:           "TokenKindPlusPlus",
	TokenKindSub:                "TokenKindSub",
	TokenKindSubAssign:          "TokenKindSubAssign",
	TokenKindSubSub:             "TokenKindSubSub",
	TokenKindMulti:              "TokenKindMulti",
	TokenKindMultiAssign:        "TokenKindMultiAssign",
	TokenKindDiv:                "TokenKindDiv",
	TokenKindDivAssign:          "TokenKindDivAssign",
	TokenKindRem:                "TokenKindRem",
	TokenKindRemAssign:          "TokenKindRemAssign",
	TokenKindBitMoveLeft:        "TokenKindBitMoveLeft",
	TokenKindBitMoveLeftAssign:  "TokenKindBitMoveLeftAssign",
	TokenKindBitMoveRight:       "TokenKindBitMoveRight",
	TokenKindBitMoveRightAssign: "TokenKindBitMoveRightAssign",
	TokenKindBitOr:              "TokenKindBitOr",
	TokenKindBitOrAssign:        "TokenKindBitOrAssign",
	TokenKindBitAnd:             "TokenKindBitAnd",
	TokenKindBitAndAssign:       "TokenKindBitAndAssign",
	TokenKindGreater:            "TokenKindGreater",
	TokenKindGreaterEquals:      "TokenKindGreaterEquals",
	TokenKindLess:               "TokenKindLess",
	TokenKindLessEquals:         "TokenKindLessEquals",
	TokenKindEquals:             "TokenKindEquals",
	TokenKindNoEquals:           "TokenKindNoEquals",
	TokenKindAnd:                "TokenKindAnd",
	TokenKindOr:                 "TokenKindOr",
	TokenKindIf:                 "if",
	TokenKindElse:               "else",
	TokenKindFor:                "for",
	TokenKindContinue:           "continue",
	TokenKindBreak:              "break",
	TokenKindFunc:               "func",
	TokenKindReturn:             "return",
	TokenKindSwitch:             "switch",
	TokenKindCase:               "case",
	TokenKindDefault:            "default",
	TokenKindImport:             "import",
	TokenKindExport:             "export",
	TokenKindInclude:            "include",
	TokenKindBase64:             "base64",
	TokenKindHex:                "hex",
}

type Token struct {
	// Kind 类型
	Kind TokenKind
	// 具体的Token值
	Value string
	// Line 开始的行号
	Line uint
	// Col 列
	Col uint
}

// IsKeyword 判断给定 kind 是否是关键字类型
func IsKeyword(kind TokenKind) bool {
	return kind >= TokenKindIf && kind < end
}

// IsComparison 判断是否是操作符类型
func IsComparison(kind TokenKind) bool {
	return kind >= TokenKindGreater && kind < TokenKindIf
}

// IsOperation 判断是否是运算符
func IsOperation(kind TokenKind) bool {
	return kind >= TokenKindPlus && kind < TokenKindGreater
}

func TokenKindToString(kind TokenKind) string {
	return tokenKindLabel[kind]
}
