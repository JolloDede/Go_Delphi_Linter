package lexer

import (
	"testing"
)

func TestLexerOperators(t *testing.T) {
	l := NewLexer("+-*/=;:")

	tok := l.next_token()
	if tok.Typ != "Token" {
		t.Errorf(`NewLexer("+-*/=;:").next_token().Typ = %q, want "Token", error`, tok.Typ)
	}
	if tok.Content != "+" {
		t.Errorf(`NewLexer("+-*/=;:").next_token().content = %q, want "+", error`, tok.Content)
	}

	tok = l.next_token()
	if tok.Typ != "Token" {
		t.Errorf(`NewLexer("+-*/=;:").next_token().Typ = %q, want "Token", error`, tok.Typ)
	}
	if tok.Content != "-" {
		t.Errorf(`NewLexer("+-*/=;:").next_token().content = %q, want "-", error`, tok.Content)
	}

	tok = l.next_token()
	if tok.Typ != "Token" {
		t.Errorf(`NewLexer("+-*/=;:").next_token().Typ = %q, want "Token", error`, tok.Typ)
	}
	if tok.Content != "*" {
		t.Errorf(`NewLexer("+-*/=;:").next_token().content = %q, want "*", error`, tok.Content)
	}

	tok = l.next_token()
	if tok.Typ != "Token" {
		t.Errorf(`NewLexer("+-*/=;:").next_token().Typ = %q, want "Token", error`, tok.Typ)
	}
	if tok.Content != "/" {
		t.Errorf(`NewLexer("+-*/=;:").next_token().content = %q, want "/", error`, tok.Content)
	}

	tok = l.next_token()
	if tok.Typ != "Token" {
		t.Errorf(`NewLexer("+-*/=;:").next_token().Typ = %q, want "Token", error`, tok.Typ)
	}
	if tok.Content != "=" {
		t.Errorf(`NewLexer("+-*/=;:").next_token().content = %q, want "=", error`, tok.Content)
	}

	tok = l.next_token()
	if tok.Typ != "Token" {
		t.Errorf(`NewLexer("+-*/=;:").next_token().Typ = %q, want "Token", error`, tok.Typ)
	}
	if tok.Content != ";" {
		t.Errorf(`NewLexer("+-*/=;:").next_token().content = %q, want ";", error`, tok.Content)
	}

	tok = l.next_token()
	if tok.Typ != "Token" {
		t.Errorf(`NewLexer("+-*/=;:").next_token().Typ = %q, want "Token", error`, tok.Typ)
	}
	if tok.Content != ":" {
		t.Errorf(`NewLexer("+-*/=;:").next_token().content = %q, want ":", error`, tok.Content)
	}
}
