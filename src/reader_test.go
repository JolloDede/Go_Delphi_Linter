package src

import (
	"testing"
)

func TestReaderPeek(t *testing.T) {
	r := NewCharReader("ab")

	c, err := r.Peek()
	if err != nil {
		t.Errorf(`NewReader("a").peek() = %q error`, err)
	}

	if c != 'a' {
		t.Errorf(`NewReader("a").peek = %q, want "a", error`, c)
	}
}

func TestReaderPeek_n(t *testing.T) {
	r := NewCharReader("ab")

	c, err := r.Peek_n(0)
	if err != nil {
		t.Errorf(`NewReader("a").peek() = %q error`, err)
	}

	if c != 'a' {
		t.Errorf(`NewReader("a").peek = %q, want "a", error`, c)
	}

	c, err = r.Peek_n(1)
	if err != nil {
		t.Errorf(`NewReader("a").peek() = %q error`, err)
	}

	if c != 'b' {
		t.Errorf(`NewReader("a").peek = %q, want "a", error`, c)
	}
}

func TestReaderNext(t *testing.T) {
	r := NewCharReader("ab")

	c, err := r.Next()
	if err != nil {
		t.Errorf(`NewReader("a").Next() = %q error`, err)
	}

	if c != 'b' {
		t.Errorf(`NewReader("a").peek = %q, want "a", error`, c)
	}
}

func TestReaderEOF(t *testing.T) {
	r := NewCharReader("a")

	if r.IsEOF() {
		t.Errorf(`NewReader("a").IsEOF = %v, want "false", error`, r.IsEOF())
	}

	r.Next()

	if !r.IsEOF() {
		t.Errorf(`NewReader("a").IsEOF = %v, want "true", error`, r.IsEOF())
	}
}

func TestReaderWhitespace(t *testing.T) {
	r := NewCharReader(`
		`)

	c, err := r.Peek_n(0)
	if err != nil {
		t.Errorf(`NewReader("a").Next() = %q error`, err)
	}

	if c != '\n' {
		t.Errorf(`NewReader("a").IsEOF = %v, want "10", error`, c)
	}

	c, err = r.Next()
	if err != nil {
		t.Errorf(`NewReader("a").Next() = %q error`, err)
	}

	if c != '\t' {
		t.Errorf(`NewReader("a").IsEOF = %v, want "10", error`, c)
	}
}

func TestReaderJump(t *testing.T) {
	r := NewCharReader("abcd")

	r.Jump(2)
	c, _ := r.Peek()

	if c != 'c' {
		t.Errorf(`NewReader("a").IsEOF = %v, want "c", error`, c)
	}

	r.Jump(1)
	c, _ = r.Peek()

	if c != 'd' {
		t.Errorf(`NewReader("a").IsEOF = %v, want "d", error`, c)
	}

}

func TestReaderCol(t *testing.T) {
	r := NewCharReader("12")

	if r.Col != 0 {
		t.Errorf(`NewReader("12").Col = %v, want 0, error`, r.Col)
	}

	r.Next()

	if r.Col != 1 {
		t.Errorf(`NewReader("12").Col = %v, want 1, error`, r.Col)
	}
}

func TestReaderRow(t *testing.T) {
	r := NewCharReader("1\n2")

	if r.Row != 0 {
		t.Errorf(`NewReader("12").Row = %v, want 0, error`, r.Row)
	}

	r.Next()

	if r.Row != 1 {
		t.Errorf(`NewReader("12").Row = %v, want 1, error`, r.Row)
	}
}
