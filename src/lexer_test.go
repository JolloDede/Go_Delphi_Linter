package src

import "testing"

func TestLexerEOF(t *testing.T) {
	l := NewLexer("")

	to := l.next_token()
	if to.Typ == "EOF" {
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
