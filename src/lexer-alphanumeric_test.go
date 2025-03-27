package src

import "testing"

func TestLexerKeywords(t *testing.T) {
	l := NewLexer("123;")
	// {"begin", "end", "unit", "interface", "implementation", "class", "record", "initialization", "finalization", "if", "then", "for, while", "do"}

	tok := l.next_token()
	if tok.Typ != "Number" {
		t.Errorf(`NewLexer("").next_token().Typ = %q, want "Number", error`, tok.Typ)
	}
	if tok.content != "123" {
		t.Errorf(`NewLexer("").next_token().content = %q, want "123;", error`, tok.content)
	}
}

func TestLexerIdentifiers(t *testing.T) {
	l := NewLexer("UMain")

	tok := l.next_token()
	if tok.Typ != "Identifier" {
		t.Errorf(`NewLexer("").next_token().Typ = %q, want "Number", error`, tok.Typ)
	}
	if tok.content != "UMain" {
		t.Errorf(`NewLexer("").next_token().content = %q, want "123;", error`, tok.content)
	}
}
