package lexer

import "testing"

func TestLexerWhitespace(t *testing.T) {
	l := NewLexer("sqlObjekt = ")

	tok := l.next_token()
	if tok.Typ != IdentifierToken {
		t.Errorf(`NewLexer("+-*/=;:").next_token().Typ = %q, want "Identifier", error`, tok.Typ)
	}
	if tok.Content != "sqlObjekt" {
		t.Errorf(`NewLexer("+-*/=;:").next_token().content = %q, want "sqlObjekt", error`, tok.Content)
	}

	// whitespace
	l.next_token()

	tok = l.next_token()
	if tok.Typ != OperatorToken {
		t.Errorf(`NewLexer("+-*/=;:").next_token().Typ = %q, want "Operator", error`, tok.Typ)
	}
	if tok.Content != "=" {
		t.Errorf(`NewLexer("+-*/=;:").next_token().content = %q, want "=", error`, tok.Content)
	}
}
