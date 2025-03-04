package src

import (
	"testing"
)

func TestReaderPeek(t *testing.T) {
	r := NewCharReader("a")

	if r.Peek() != 'a' {
		t.Errorf(`NewReader("a").peek = %q, want "a", error`, r.Peek())
	}
}

func TestReaderNext(t *testing.T) {
	r := NewCharReader("abc")

	if r.Next() != 'a' {
		t.Errorf(`NewReader("a").peek = %q, want "a", error`, r.Peek())
	}
	if r.Next() != 'b' {
		t.Errorf(`NewReader("a").peek = %q, want "b", error`, r.Peek())
	}
	if r.Next() != 'c' {
		t.Errorf(`NewReader("a").peek = %q, want "c", error`, r.Peek())
	}
}

func TestReaderEOF(t *testing.T) {
	r := NewCharReader("a")

	if r.IsEOF() {
		t.Errorf(`NewReader("a").IsEOF = %v, want "false", error`, r.IsEOF())
	}

	r.Next()

	if !r.IsEOF() {
		t.Errorf(`NewReader("a").peek = %q, want "true", error`, r.Peek())
	}
}
