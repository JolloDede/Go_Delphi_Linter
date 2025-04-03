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
	l := NewLexer(`begin end unit interface implementation class record initialization finalization if then for while do `)

	tok := l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "begin" {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().content = %q, want "begin", error`, tok.content)
	}

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "end" {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().content = %q, want "end", error`, tok.content)
	}

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "unit" {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().content = %q, want "unit", error`, tok.content)
	}

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "interface" {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().content = %q, want "interface", error`, tok.content)
	}

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "implementation" {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().content = %q, want "implementation", error`, tok.content)
	}

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "class" {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().content = %q, want "class", error`, tok.content)
	}

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "record" {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().content = %q, want "record", error`, tok.content)
	}

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "initialization" {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().content = %q, want "initialization", error`, tok.content)
	}

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "finalization" {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().content = %q, want "finalization", error`, tok.content)
	}

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "if" {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().content = %q, want "if", error`, tok.content)
	}

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "then" {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().content = %q, want "then", error`, tok.content)
	}

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "for" {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().content = %q, want "for", error`, tok.content)
	}

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "while" {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().content = %q, want "while", error`, tok.content)
	}

	tok = l.next_token()
	if tok.Typ != KeywordToken {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().Typ = %q, want "Keyword", error`, tok.Typ)
	}
	if tok.content != "do" {
		t.Errorf(`NewLexer("begin end unit interface implementation class record initialization finalization if then for while do ").next_token().content = %q, want "do", error`, tok.content)
	}
}
