package main

import (
	"fmt"
	"os"
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
		// 字句エラー
		// 本では Lox.error を呼んでいるが、依存関係がおかしい気がするので一旦エラーログを雑に吐く
		fmt.Printf("Unexpected character: %d\n", s.Line)
		return fmt.Errorf("unexpected character: %d", s.Line)
	}
	return nil
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
