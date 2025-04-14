package lexer

import (
	"testing"
)

func TestLexerIdentifiers(t *testing.T) {
	l := NewLexer("UMain ")

	tok := l.next_token()
	if tok.Typ != "Identifier" {
		t.Errorf(`NewLexer("UMain").next_token().Typ = %q, want "Identifier", error`, tok.Typ)
	}
	if tok.Content != "UMain" {
		t.Errorf(`NewLexer("UMain").next_token().content = %q, want "UMain", error`, tok.Content)
	}
}

func TestLexerKeywords(t *testing.T) {
	l := NewLexer("begin ")

	tok := l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("begin ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.Content != "begin" {
		t.Errorf(`NewLexer("begin ").next_token().content = %q, want "begin", error`, tok.Content)
	}

	l = NewLexer("end ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("end ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.Content != "end" {
		t.Errorf(`NewLexer("end ").next_token().content = %q, want "end", error`, tok.Content)
	}

	l = NewLexer("unit ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("unit ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.Content != "unit" {
		t.Errorf(`NewLexer("unit ").next_token().content = %q, want "unit", error`, tok.Content)
	}

	l = NewLexer("interface ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("interface ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.Content != "interface" {
		t.Errorf(`NewLexer("interface ").next_token().content = %q, want "interface", error`, tok.Content)
	}

	l = NewLexer("implementation ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("implementation ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.Content != "implementation" {
		t.Errorf(`NewLexer("implementation ").next_token().content = %q, want "implementation", error`, tok.Content)
	}

	l = NewLexer("class ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("class ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.Content != "class" {
		t.Errorf(`NewLexer("class ").next_token().content = %q, want "class", error`, tok.Content)
	}

	l = NewLexer("record ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("record ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.Content != "record" {
		t.Errorf(`NewLexer("record ").next_token().content = %q, want "record", error`, tok.Content)
	}

	l = NewLexer("initialization ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("initialization ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.Content != "initialization" {
		t.Errorf(`NewLexer("initialization ").next_token().content = %q, want "initialization", error`, tok.Content)
	}

	l = NewLexer("finalization ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("finalization ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.Content != "finalization" {
		t.Errorf(`NewLexer("finalization ").next_token().content = %q, want "finalization", error`, tok.Content)
	}

	l = NewLexer("if ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("if ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.Content != "if" {
		t.Errorf(`NewLexer("if ").next_token().content = %q, want "if", error`, tok.Content)
	}

	l = NewLexer("then ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("then ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.Content != "then" {
		t.Errorf(`NewLexer("then ").next_token().content = %q, want "then", error`, tok.Content)
	}

	l = NewLexer("for ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("for ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.Content != "for" {
		t.Errorf(`NewLexer("for ").next_token().content = %q, want "for", error`, tok.Content)
	}

	l = NewLexer("while ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("while ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.Content != "while" {
		t.Errorf(`NewLexer("while ").next_token().content = %q, want "while", error`, tok.Content)
	}

	l = NewLexer("do ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("do ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.Content != "do" {
		t.Errorf(`NewLexer("do ").next_token().content = %q, want "do", error`, tok.Content)
	}
}
