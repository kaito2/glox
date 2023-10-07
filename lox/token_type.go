package lox

import "fmt"

type TokenType string

const (
	// 記号1個のトークン
	LeftParenTokenType  TokenType = "left-paren"
	RightParenTokenType TokenType = "right-paren"
	LeftBraceTokenType  TokenType = "left-brace"
	RightBraceTokenType TokenType = "right-brace"
	CommaTokenType      TokenType = "comma"
	DotTokenType        TokenType = "dot"
	MinusTokenType      TokenType = "minus"
	PlusTokenType       TokenType = "plus"
	SemicolonTokenType  TokenType = "semicolon"
	SlashTokenType      TokenType = "slash"
	StarTokenType       TokenType = "star"

	// 記号1個または2個によるトークン
	BangTokenType         TokenType = "bang"
	BangEqualTokenType    TokenType = "bang-equal"
	EqualTokenType        TokenType = "equal"
	EqualEqualTokenType   TokenType = "equal-equal"
	GreaterTokenType      TokenType = "greater"
	GreaterEqualTokenType TokenType = "greater-equal"
	LessTokenType         TokenType = "less"
	LessEqualTokenType    TokenType = "less-equal"

	// リテラル
	IdentifierTokenType TokenType = "identifier"
	StringTokenType     TokenType = "string"
	NumberTokenType     TokenType = "number"

	// キーワード
	AndTokenType    TokenType = "and"
	ClassTokenType  TokenType = "class"
	ElseTokenType   TokenType = "else"
	FalseTokenType  TokenType = "false"
	FunTokenType    TokenType = "fun"
	ForTokenType    TokenType = "for"
	IfTokenType     TokenType = "if"
	NilTokenType    TokenType = "nil"
	OrTokenType     TokenType = "or"
	PrintTokenType  TokenType = "print"
	ReturnTokenType TokenType = "return"
	SuperTokenType  TokenType = "super"
	ThisTokenType   TokenType = "this"
	TrueTokenType   TokenType = "true"
	VarTokenType    TokenType = "var"
	WhileTokenType  TokenType = "while"

	// その他
	EOFTokenType TokenType = "eof"
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}

func NewToken(
	tokenType TokenType,
	lexeme string,
	literal any,
	line int,
) Token {
	return Token{
		Type:    tokenType,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s %+v", t.Type, t.Lexeme, t.Literal)
}
