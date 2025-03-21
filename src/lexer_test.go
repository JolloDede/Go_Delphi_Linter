package src

import (
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
	l := NewLexer("'eggs'")

	tok := l.next_token()
	if tok.Typ != "String" {
		t.Errorf(`NewLexer("").next_token().Typ = %q, want "String", error`, tok.Typ)
	}
	if tok.content != "eggs" {
		t.Errorf(`NewLexer("").next_token().content = %q, want "eggs", error`, tok.content)
	}
}

// func TestLexerStringLiteralMulitline(t *testing.T) {
// 	panic("Not yet implemented")

// 	l := NewLexer(`'''
// 	mulitline
// 	'''`)

// 	tok := l.next_token()
// 	if tok.Typ != "String" {
// 		t.Errorf(`NewLexer("").next_token().Typ = %q, want "String", error`, tok.Typ)
// 	}

// 	fmt.Println(tok.content)
// 	if tok.content != "mulitline\n\t" {
// 		t.Errorf(`NewLexer("").next_token().content = %q, want "\tmulitline\n\t", error`, tok.content)
// 	}
// }

func TestLexerComment(t *testing.T) {
	l := NewLexer("// Comment\n")

	tok := l.next_token()
	if tok.Typ != "Comment" {
		t.Errorf(`NewLexer("").next_token().Typ = %q, want "Comment", error`, tok.Typ)
	}
	if tok.content != " Comment" {
		t.Errorf(`NewLexer("").next_token().content = %q, want " Comment", error`, tok.content)
	}
}

func TestLexerCommentMultiline(t *testing.T) {
	l := NewLexer("{Comment\nNextLine}")

	tok := l.next_token()
	if tok.Typ != "Comment" {
		t.Errorf(`NewLexer("").next_token().Typ = %q, want "Comment", error`, tok.Typ)
	}
	if tok.content != "Comment\nNextLine" {
		t.Errorf(`NewLexer("").next_token().content = %q, want "Comment\nNextLine", error`, tok.content)
	}
}

func TestLexerConditionalCompilation(t *testing.T) {
	l := NewLexer("{$DEFINE DEBUG}")

	tok := l.next_token()
	if tok.Typ != "CondComp" {
		t.Errorf(`NewLexer("").next_token().Typ = %q, want "CondComp", error`, tok.Typ)
	}
	if tok.content != "$DEFINE DEBUG" {
		t.Errorf(`NewLexer("").next_token().content = %q, want "$DEFINE DEBUG", error`, tok.content)
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
