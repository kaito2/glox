package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Scanner(t *testing.T) {
	testCases := []struct {
		name      string
		given     string
		wantError bool
		want      []Token
	}{
		{"正常系1", "()", false, []Token{
			{Type: LeftParenTokenType, Lexeme: "("},
			{Type: RightParenTokenType, Lexeme: ")"},
			{Type: EOFTokenType},
		}},
		{"正常系2", "( )", false, []Token{
			{Type: LeftParenTokenType, Lexeme: "("},
			{Type: RightParenTokenType, Lexeme: ")"},
			{Type: EOFTokenType},
		}},
		{"正常系3", "(( )){}", false, []Token{
			{Type: LeftParenTokenType, Lexeme: "("},
			{Type: LeftParenTokenType, Lexeme: "("},
			{Type: RightParenTokenType, Lexeme: ")"},
			{Type: RightParenTokenType, Lexeme: ")"},
			{Type: LeftBraceTokenType, Lexeme: "{"},
			{Type: RightBraceTokenType, Lexeme: "}"},
			{Type: EOFTokenType},
		}},
		{"文字列リテラル", `"hello world"`, false, []Token{
			{Type: StringTokenType, Lexeme: `"hello world"`, Literal: "hello world"},
			{Type: EOFTokenType},
		}},
	}

	for _, testCase := range testCases {
		sut := NewScanner(testCase.given)
		got, err := sut.ScanTokens()
		if testCase.wantError && err == nil {
			t.Fatal("This test case expect error, but err is nil.")
		} else if !testCase.wantError && err != nil {
			t.Fatal("This test case don't expected error, but err is not nil.")
		}
		if diff := cmp.Diff(testCase.want, got); diff != "" {
			t.Fatal(diff)
		}
	}
}
