package src

import (
	"testing"
)

func TestReaderPeek(t *testing.T) {
	r := NewCharReader("ab")

	c := r.Peek()
	if c != 'a' {
		t.Errorf(`NewReader("a").peek() = %q, want "a", error`, c)
	}

	if c != 'a' {
		t.Errorf(`NewReader("a").peek() = %q, want "b", error`, c)
	}
}

func TestReaderPeekPanik(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(`NewReader("").peek() panic was expected`)
		}
	}()

	r := NewCharReader("")

	// Call the function that will panic
	r.Peek()
	// If the panic was not caught, the test will fail
	t.Errorf(`NewReader("").peek() panic was expected`)
}

func TestReaderPeek_n(t *testing.T) {
	r := NewCharReader("ab")

	c := r.Peek_n(0)

	if c != 'a' {
		t.Errorf(`NewReader("a").peek = %q, want "a", error`, c)
	}

	c = r.Peek_n(1)

	if c != 'b' {
		t.Errorf(`NewReader("a").peek = %q, want "a", error`, c)
	}
}

func TestReaderPeek_nPanik(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(`NewReader("").Peek_n(3) panic was expected`)
		}
	}()

	r := NewCharReader("12")

	// Call the function that will panic
	r.Peek_n(3)
	// If the panic was not caught, the test will fail
	t.Errorf(`NewReader("").Peek_n(3) panic was expected`)
}

func TestReaderNext(t *testing.T) {
	r := NewCharReader("ab")

	c := r.Next()

	if c != 'b' {
		t.Errorf(`NewReader("ab").peek = %q, want "b", error`, c)
	}
}

func TestReaderNextPanik(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(`NewCharReader("1").Next() panic was expected`)
		}
	}()

	r := NewCharReader("1")

	// Call the function that will panic
	r.Next()
	// If the panic was not caught, the test will fail
	t.Errorf(`NewCharReader("1").Next(3) panic was expected`)
}

func TestReaderEOF(t *testing.T) {
	r := NewCharReader("ab")

	if r.IsEOF() {
		t.Errorf(`NewReader("ab").IsEOF = %v, want "false", error`, r.IsEOF())
	}

	r.Next()

	if !r.IsEOF() {
		t.Errorf(`NewReader("ab").IsEOF = %v, want "true", error`, r.IsEOF())
	}
}

func TestReaderWhitespace(t *testing.T) {
	r := NewCharReader(`
		`)

	c := r.Peek_n(0)

	if c != '\n' {
		t.Errorf(`NewReader("a").IsEOF = %v, want "10", error`, c)
	}

	c = r.Next()

	if c != '\t' {
		t.Errorf(`NewReader("a").IsEOF = %v, want "10", error`, c)
	}
}

func TestReaderJump(t *testing.T) {
	r := NewCharReader("abcd")

	r.Jump(2)
	c := r.Peek()

	if c != 'c' {
		t.Errorf(`NewReader("a").IsEOF = %v, want "c", error`, c)
	}

	r.Jump(1)
	c = r.Peek()

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
