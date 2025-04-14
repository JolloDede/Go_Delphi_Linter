package lexer

import "testing"

func TestLexerSql(t *testing.T) {
	l := NewLexer("sqlObjekt = 'SELECT * from tObject';")

	tok := l.next_token()
	if tok.Typ != IdentifierToken {
		t.Errorf(`NewLexer("sqlObjekt = 'SELECT * from tObject';").next_token().Typ = %q, want "Identifier", error`, tok.Typ)
	}
	if tok.Content != "sqlObjekt" {
		t.Errorf(`NewLexer("sqlObjekt = 'SELECT * from tObject';").next_token().content = %q, want "sqlObjekt", error`, tok.Content)
	}

	// whitespace
	l.next_token()

	tok = l.next_token()
	if tok.Typ != OperatorToken {
		t.Errorf(`NewLexer("sqlObjekt = 'SELECT * from tObject';").next_token().Typ = %q, want "Operator", error`, tok.Typ)
	}
	if tok.Content != "=" {
		t.Errorf(`NewLexer("sqlObjekt = 'SELECT * from tObject';").next_token().content = %q, want "=", error`, tok.Content)
	}

	// whitespace
	l.next_token()

	tok = l.next_token()
	if tok.Typ != StringToken {
		t.Errorf(`NewLexer("sqlObjekt = 'SELECT * from tObject';").next_token().Typ = %q, want "String", error`, tok.Typ)
	}
	if tok.Content != "SELECT * from tObject" {
		t.Errorf(`NewLexer("sqlObjekt = 'SELECT * from tObject';").next_token().content = %q, want "SELECT * from tObject", error`, tok.Content)
	}

	tok = l.next_token()
	if tok.Typ != OperatorToken {
		t.Errorf(`NewLexer("sqlObjekt = 'SELECT * from tObject';").next_token().Typ = %q, want "Operator", error`, tok.Typ)
	}
	if tok.Content != ";" {
		t.Errorf(`NewLexer("sqlObjekt = 'SELECT * from tObject';").next_token().content = %q, want ";", error`, tok.Content)
	}
}
