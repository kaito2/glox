package lox

import (
	"fmt"
	"os"
	"strconv"
)

type Scanner struct {
	Source string
	Tokens []Token

	// TODO: Private でいいかも?
	Start   int
	Current int
	Line    int
}

func NewScanner(
	source string,
) Scanner {
	return Scanner{
		Source: source,
		Tokens: []Token{},
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.Current >= len(s.Source)
}

func (s *Scanner) hasNext() bool {
	return s.Current < len(s.Source)
}

func (s *Scanner) ScanTokens() ([]Token, error) {
	for !s.isAtEnd() {
		s.Start = s.Current
		if err := s.ScanToken(); err != nil {
			return nil, fmt.Errorf("failed to scan token: %w", err)
		}
	}

	// 末尾には EOF トークンを挿入
	s.Tokens = append(s.Tokens, NewToken(EOFTokenType, "", nil, s.Line))
	return s.Tokens, nil
}

func (s *Scanner) ScanToken() error {
	c := s.advance()
	switch c {
	case '(':
		s.addTokenWithType(LeftParenTokenType)
	case ')':
		s.addTokenWithType(RightParenTokenType)
	case '{':
		s.addTokenWithType(LeftBraceTokenType)
	case '}':
		s.addTokenWithType(RightBraceTokenType)
	case ',':
		s.addTokenWithType(CommaTokenType)
	case '.':
		s.addTokenWithType(DotTokenType)
	case '-':
		s.addTokenWithType(MinusTokenType)
	case '+':
		s.addTokenWithType(PlusTokenType)
	case ';':
		s.addTokenWithType(SemicolonTokenType)
	case '*':
		s.addTokenWithType(StarTokenType)
	case '!':
		if s.match('=') {
			s.addTokenWithType(BangEqualTokenType)
		} else {
			s.addTokenWithType(BangTokenType)
		}
	case '=':
		if s.match('=') {
			s.addTokenWithType(EqualEqualTokenType)
		} else {
			s.addTokenWithType(EqualTokenType)
		}
	case '<':
		if s.match('=') {
			s.addTokenWithType(LessEqualTokenType)
		} else {
			s.addTokenWithType(LessTokenType)
		}
	case '>':
		if s.match('=') {
			s.addTokenWithType(GreaterEqualTokenType)
		} else {
			s.addTokenWithType(GreaterTokenType)
		}
	case '/':
		if s.match('/') {
			// コメントは行の末尾まで続く
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addTokenWithType(SlashTokenType)
		}
	case ' ':
	case '\r':
	case '\t':
		break
	case '\n':
		s.Line++
	case '"':
		s.string()
	default:
		if isDigit(c) {
			s.number()
			break
		} else if isAlpha(c) {
			s.identifier()
			break
		}
		// 字句エラー
		// 本では Lox.error を呼んでいるが、依存関係がおかしい気がするので一旦エラーログを雑に吐く
		fmt.Printf("Unexpected character: %d\n", s.Line)
		return fmt.Errorf("unexpected character: %d", s.Line)
	}
	return nil
}

func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := s.Source[s.Start:s.Current]
	tokenType, ok := keywordType(text)
	if !ok {
		tokenType = IdentifierTokenType
	}
	s.addTokenWithType(tokenType)
}

func keywordType(text string) (TokenType, bool) {
	switch text {
	case "and":
		return AndTokenType, true
	case "class":
		return ClassTokenType, true
	case "else":
		return ElseTokenType, true
	case "false":
		return FalseTokenType, true
	case "for":
		return ForTokenType, true
	case "fun":
		return FunTokenType, true
	case "if":
		return IfTokenType, true
	case "nil":
		return NilTokenType, true
	case "or":
		return OrTokenType, true
	case "print":
		return PrintTokenType, true
	case "return":
		return ReturnTokenType, true
	case "super":
		return SuperTokenType, true
	case "this":
		return ThisTokenType, true
	case "true":
		return TrueTokenType, true
	case "var":
		return VarTokenType, true
	case "while":
		return WhileTokenType, true
	default:
		return UnknownTokenType, false
	}
}

func isAlphaNumeric(c rune) bool {
	return isDigit(c) || isAlpha(c)
}

func isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}

	// 小数部があるか確認
	if s.peek() == '.' && isDigit(s.peekNext()) {
		// 小数点を消費
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}
	}

	literal, err := strconv.ParseFloat(string(s.Source[s.Start:s.Current]), 64)
	if err != nil {
		panic("Failed to strconv.ParseFloat") // TODO: Don't panic.
	}
	s.addToken(NumberTokenType, literal)
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.Line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		// 本では Lox.error を呼んでいるが、依存関係がおかしい気がするので一旦エラーログを雑に吐く
		fmt.Printf("Unexpected character: %d\n", s.Line)
		os.Exit(65) // TODO: Return error
	}

	// 右の引用符を消費
	s.advance()

	value := s.Source[s.Start+1 : s.Current-1] // TODO: Validate index.
	s.addToken(StringTokenType, value)
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		// TODO: '\0' のかわりに使ったが、あってる?
		return rune(0)
	}
	return rune(s.Source[s.Current])
}

func (s *Scanner) peekNext() rune {
	if s.Current+1 >= len(s.Source) {
		// TODO: '\0' のかわりに使ったが、あってる?
		return rune(0)
	}
	return rune(s.Source[s.Current+1])
}

// 条件付き advance
func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}
	if rune(s.Source[s.Current]) == expected {
		s.Current++
		return true
	}

	return false
}

func (s *Scanner) advance() rune {
	r := rune(s.Source[s.Current])
	s.Current++
	return r
}

// NOTE: Go はオーバーロードできないので WithType suffix をつける
func (s *Scanner) addTokenWithType(tokenType TokenType) {
	s.addToken(tokenType, nil)
}

func (s *Scanner) addToken(tokenType TokenType, literal any) {
	// TODO: s.Start, s.Current のバリデーション
	// こんなに状態をもつ必要あるか?
	text := s.Source[s.Start:s.Current]
	s.Tokens = append(s.Tokens, NewToken(tokenType, text, literal, s.Line))
}
