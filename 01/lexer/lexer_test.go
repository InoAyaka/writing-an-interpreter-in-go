package lexer

import (
	"Book_waiig/01/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
		let five = 5;
		let ten = 10;

		let add = fn(x,y){
			x + y;
		};

		let result = add(five,ten);
		!-/*5;
		5 < 10 > 5;

		if (5 < 10){
			return true;
		}else{
			return false;
		}
	`

	tests := []struct {
		expextedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDNET, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDNET, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDNET, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDNET, "x"},
		{token.COMMA, ","},
		{token.IDNET, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDNET, "x"},
		{token.PLUS, "+"},
		{token.IDNET, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDNET, "result"},
		{token.ASSIGN, "="},
		{token.IDNET, "add"},
		{token.LPAREN, "("},
		{token.IDNET, "five"},
		{token.COMMA, ","},
		{token.IDNET, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expextedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expextedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenliteral wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
