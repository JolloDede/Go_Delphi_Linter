package src

import "testing"

func TestLexerIdentifiers(t *testing.T) {
	l := NewLexer("UMain ")

	tok := l.next_token()
	if tok.Typ != "Identifier" {
		t.Errorf(`NewLexer("UMain").next_token().Typ = %q, want "Identifier", error`, tok.Typ)
	}
	if tok.content != "UMain" {
		t.Errorf(`NewLexer("UMain").next_token().content = %q, want "UMain", error`, tok.content)
	}
}

func TestLexerKeywords(t *testing.T) {
	l := NewLexer("begin ")

	tok := l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("begin ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "begin" {
		t.Errorf(`NewLexer("begin ").next_token().content = %q, want "begin", error`, tok.content)
	}

	l = NewLexer("end ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("end ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "end" {
		t.Errorf(`NewLexer("end ").next_token().content = %q, want "end", error`, tok.content)
	}

	l = NewLexer("unit ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("unit ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "unit" {
		t.Errorf(`NewLexer("unit ").next_token().content = %q, want "unit", error`, tok.content)
	}

	l = NewLexer("interface ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("interface ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "interface" {
		t.Errorf(`NewLexer("interface ").next_token().content = %q, want "interface", error`, tok.content)
	}

	l = NewLexer("implementation ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("implementation ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "implementation" {
		t.Errorf(`NewLexer("implementation ").next_token().content = %q, want "implementation", error`, tok.content)
	}

	l = NewLexer("class ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("class ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "class" {
		t.Errorf(`NewLexer("class ").next_token().content = %q, want "class", error`, tok.content)
	}

	l = NewLexer("record ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("record ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "record" {
		t.Errorf(`NewLexer("record ").next_token().content = %q, want "record", error`, tok.content)
	}

	l = NewLexer("initialization ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("initialization ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "initialization" {
		t.Errorf(`NewLexer("initialization ").next_token().content = %q, want "initialization", error`, tok.content)
	}

	l = NewLexer("finalization ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("finalization ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "finalization" {
		t.Errorf(`NewLexer("finalization ").next_token().content = %q, want "finalization", error`, tok.content)
	}

	l = NewLexer("if ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("if ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "if" {
		t.Errorf(`NewLexer("if ").next_token().content = %q, want "if", error`, tok.content)
	}

	l = NewLexer("then ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("then ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "then" {
		t.Errorf(`NewLexer("then ").next_token().content = %q, want "then", error`, tok.content)
	}

	l = NewLexer("for ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("for ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "for" {
		t.Errorf(`NewLexer("for ").next_token().content = %q, want "for", error`, tok.content)
	}

	l = NewLexer("while ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("while ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "while" {
		t.Errorf(`NewLexer("while ").next_token().content = %q, want "while", error`, tok.content)
	}

	l = NewLexer("do ")

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("do ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "do" {
		t.Errorf(`NewLexer("do ").next_token().content = %q, want "do", error`, tok.content)
	}
}
