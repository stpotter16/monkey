package lexer_test

import (
	"monkey/internal/lexer"
	"monkey/internal/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{
			expectedType:    token.ASSIGN,
			expectedLiteral: "=",
		},
		{
			expectedType:    token.PLUS,
			expectedLiteral: "+",
		},
		{
			expectedType:    token.LPAREN,
			expectedLiteral: "(",
		},
		{
			expectedType:    token.RPAREN,
			expectedLiteral: ")",
		},
		{
			expectedType:    token.LBRACKET,
			expectedLiteral: "{",
		},
		{
			expectedType:    token.RBRACKET,
			expectedLiteral: "}",
		},
		{
			expectedType:    token.COMMA,
			expectedLiteral: ",",
		},
		{
			expectedType:    token.SEMICOLON,
			expectedLiteral: ";",
		},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q got=%q", i, tok.Type, tt.expectedType)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - token literal wrong. expected=%q got=%q", i, tok.Literal, tt.expectedLiteral)
		}
	}
}
