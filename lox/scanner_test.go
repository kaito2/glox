package lox

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
		{
			"数値リテラル", "1234", false, []Token{
				{Type: NumberTokenType, Lexeme: "1234", Literal: float64(1234)},
				{Type: EOFTokenType},
			},
		},
	}

	for _, testCase := range testCases {
		sut := NewScanner(testCase.given)
		got, err := sut.ScanTokens()
		if testCase.wantError && err == nil {
			t.Fatal("This test case expect error, but err is nil.")
		} else if !testCase.wantError && err != nil {
			t.Fatal("This test case don't expected error, but err is not nil: ", err)
		}
		if diff := cmp.Diff(testCase.want, got); diff != "" {
			t.Fatal(diff)
		}
	}
}

func Test_isDigit(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"a", args{'a'}, false},
		{"0", args{'0'}, true},
		{"1", args{'1'}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDigit(tt.args.c); got != tt.want {
				t.Errorf("isDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
