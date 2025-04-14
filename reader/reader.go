package reader

type EOFError struct{}

func (e *EOFError) Error() string {
	return "EOF error"
}

type OOBError struct{}

func (e *OOBError) Error() string {
	return "Out of Bounds error"
}

type CharReader struct {
	content []rune
	index   int
	Row     int
	Col     int
}

func NewCharReader(input string) CharReader {
	return CharReader{content: []rune(input), Row: 0, Col: 0, index: 0}
}

func (r *CharReader) Peek() rune {
	if r.index > len(r.content)-1 {
		panic("Peek EOF")
	}

	return r.content[r.index]
}

func (r *CharReader) Peek_n(i int) rune {
	if r.index+i > len(r.content)-1 {
		panic("Peek_n EOF")
	}
	return r.content[r.index+i]
}

func (r *CharReader) Next() rune {
	r.index++
	r.Col++

	if r.index >= len(r.content) {
		panic("Next EOF")
	}

	if r.content[r.index] == '\n' {
		r.Row++
		r.Col = 0
	}

	return r.content[r.index]
}

func (r *CharReader) Next_n(n int) rune {

	for i := 0; i > n; i++ {
		r.Next()
	}

	return r.content[r.index]
}

func (r *CharReader) Jump(i int) {
	r.index += i
}

func (r *CharReader) IsEOF() bool {
	return r.index > len(r.content)-1
}
