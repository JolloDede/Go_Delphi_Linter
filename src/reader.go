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

	if r.index >= len(r.content) {
		return ' ', &EOFError{}
	}

	return r.content[r.index], nil
}

func (r *CharReader) Next_n(i int) (byte, error) {
	r.index += i

	if r.index >= len(r.content) {
		return ' ', &EOFError{}
	}

	return r.content[r.index], nil
}

func (r *CharReader) Jump(i int) {
	r.index += i
}

func (r *CharReader) IsEOF() bool {
	return r.index >= len(r.content)
}
