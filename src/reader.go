package src

type EOFError struct{}

func (e *EOFError) Error() string {
	return "EOF error"
}

type OOBError struct{}

func (e *OOBError) Error() string {
	return "Out of Bounds error"
}

type CharReader struct {
	content string
	index   int
	Row     int
	Col     int
}

func NewCharReader(input string) CharReader {
	return CharReader{content: input, Row: 0, Col: 0, index: 0}
}

func (r *CharReader) Peek() (byte, error) {
	if r.index > len(r.content)-1 {
		return ' ', &EOFError{}
	}
	return r.content[r.index], nil
}

func (r *CharReader) Peek_n(i int) (byte, error) {
	if i < 0 {
		return ' ', &OOBError{}
	}
	if i > len(r.content)-1 {
		return ' ', &EOFError{}
	}
	return r.content[i], nil
}

func (r *CharReader) Next() (byte, error) {
	r.index++
	r.Col++

	if r.index >= len(r.content) {
		return ' ', &EOFError{}
	}

	if r.content[r.index] == '\n' {
		r.Row++
		r.Col = 0
	}

	return r.content[r.index], nil
}

func (r *CharReader) Next_n(n int) (byte, error) {

	for i := 0; i > n; i++ {
		r.Next()
	}

	return r.content[r.index], nil
}

func (r *CharReader) Jump(i int) {
	r.index += i
}

func (r *CharReader) IsEOF() bool {
	return r.index >= len(r.content)
}
