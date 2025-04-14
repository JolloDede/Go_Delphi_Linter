package lexer

import (
	"fmt"
	"testing"
)

func TestLexerEOF(t *testing.T) {
	l := NewLexer("")

	to := l.next_token()
	if to.Typ != "EOF" {
		t.Errorf(`NewLexer("").next_token = %q, want "Token{typ: "EOF"}", error`, to)
	}
}

func TestLexerUnitDef(t *testing.T) {
	// l := NewLexer(`unit uMain;
	// 	`)

	// to := l.next_token()
	// if to.Typ == "EOF" {
	// 	t.Errorf(`NewLexer("").next_token = %q, want "Token{typ: "EOF"}", error`, to)
	// }
}

func TestLexerStringLiteral(t *testing.T) {
	l := NewLexer("'eggs' ")

	tok := l.next_token()
	if tok.Typ != "String" {
		t.Errorf(`NewLexer("").next_token().Typ = %q, want "String", error`, tok.Typ)
	}
	if tok.Content != "eggs" {
		t.Errorf(`NewLexer("").next_token().content = %q, want "eggs", error`, tok.Content)
	}
}

func TestLexerStringLiteralEscape(t *testing.T) {
	l := NewLexer("'test''s' ")

	tok := l.next_token()
	if tok.Typ != "String" {
		t.Errorf(`NewLexer("").next_token().Typ = %q, want "String", error`, tok.Typ)
	}
	if tok.Content != "test's" {
		t.Errorf(`NewLexer("").next_token().content = %q, want "test's", error`, tok.Content)
	}
}

func TestLexerStringLiteralMulitline(t *testing.T) {
	l := NewLexer(`'''
	mulitline
	''' `)

	tok := l.next_token()
	if tok.Typ != "String" {
		t.Errorf(`NewLexer("").next_token().Typ = %q, want "String", error`, tok.Typ)
	}

	fmt.Println(tok.Content)
	if tok.Content != "\tmulitline\n\t" {
		t.Errorf(`NewLexer("").next_token().content = %q, want "\tmulitline\n\t", error`, tok.Content)
	}
}

func TestLexerStringLiteralMulitlineEscape(t *testing.T) {
	l := NewLexer(`'''
	mulitline''s
	''' `)

	tok := l.next_token()
	if tok.Typ != "String" {
		t.Errorf(`NewLexer("").next_token().Typ = %q, want "String", error`, tok.Typ)
	}

	fmt.Println(tok.Content)
	if tok.Content != "\tmulitline's\n\t" {
		t.Errorf(`NewLexer("").next_token().content = %q, want "\tmulitline's\n\t", error`, tok.Content)
	}
}

func TestLexerComment(t *testing.T) {
	l := NewLexer("// Comment\n")

	tok := l.next_token()
	if tok.Typ != "Comment" {
		t.Errorf(`NewLexer("// Comment\n").next_token().Typ = %q, want "Comment", error`, tok.Typ)
	}
	if tok.Content != " Comment" {
		t.Errorf(`NewLexer("// Comment\n").next_token().content = %q, want " Comment", error`, tok.Content)
	}
}

func TestLexerCommentMultiline(t *testing.T) {
	l := NewLexer("{Comment\nNextLine}")

	tok := l.next_token()
	if tok.Typ != "Comment" {
		t.Errorf(`NewLexer("{Comment\nNextLine}").next_token().Typ = %q, want "Comment", error`, tok.Typ)
	}
	if tok.Content != "Comment\nNextLine" {
		t.Errorf(`NewLexer("{Comment\nNextLine}").next_token().content = %q, want "Comment\nNextLine", error`, tok.Content)
	}
}

func TestLexerConditionalCompilation(t *testing.T) {
	l := NewLexer("{$DEFINE DEBUG}")

	tok := l.next_token()
	if tok.Typ != "CondComp" {
		t.Errorf(`NewLexer("").next_token().Typ = %q, want "CondComp", error`, tok.Typ)
	}
	if tok.Content != "$DEFINE DEBUG" {
		t.Errorf(`NewLexer("").next_token().content = %q, want "$DEFINE DEBUG", error`, tok.Content)
	}
}

func TestLexerConditionalCompilationEOF(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(`NewLexer("").next_token() panic was expected`)
		}
	}()

	l := NewLexer("{$DEFINE DEBUG")

	// Call the function that will panic
	l.next_token()
	// If the panic was not caught, the test will fail
	t.Errorf(`NewLexer("").next_token() panic was expected`)
}

func TestLexerNumber(t *testing.T) {
	l := NewLexer("123;")

	tok := l.next_token()
	if tok.Typ != "Number" {
		t.Errorf(`NewLexer("").next_token().Typ = %q, want "Number", error`, tok.Typ)
	}
	if tok.Content != "123" {
		t.Errorf(`NewLexer("").next_token().content = %q, want "123;", error`, tok.Content)
	}
}

func TestLexerFloatNumber(t *testing.T) {
	l := NewLexer("123.123;")

	tok := l.next_token()
	if tok.Typ != "Number" {
		t.Errorf(`NewLexer("").next_token().Typ = %q, want "Number", error`, tok.Typ)
	}
	if tok.Content != "123.123" {
		t.Errorf(`NewLexer("").next_token().content = %q, want "123.123", error`, tok.Content)
	}
}
