package parser

import (
	"Book_waiig/02/ast"
	"Book_waiig/02/lexer"
	"testing"
)

func TestLetStatment(t *testing.T) {
	input := `
	let x  5;
	let = 10;
	let 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParserProgram()
	checkParseErrors(t, p)
	if program == nil {
		t.Fatal("ParserProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements dose not contain 3 statemets. got = %d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatment(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatment(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letstmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letstmt.Name.Value != name {
		t.Errorf("letstmt.Name.Value not '%s'. got=%s", name, letstmt.Name.Value)
		return false
	}

	if letstmt.Name.TokenLiteral() != name {
		t.Errorf("letstmt.Name.TokenLiteral() not '%s'. got=%s", name, letstmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
