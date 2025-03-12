package src

import (
	"testing"
)

func TestReaderPeek(t *testing.T) {
	r := NewCharReader("ab")

	c, err := r.Peek(0)
	if err != nil {
		t.Errorf(`NewReader("a").peek() = %q error`, err)
	}

	if c != 'a' {
		t.Errorf(`NewReader("a").peek = %q, want "a", error`, c)
	}

	c, err = r.Peek(1)
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

	c, err := r.Peek(0)
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
