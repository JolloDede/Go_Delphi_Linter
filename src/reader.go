package src

type CharReader struct {
	content string
	i       int
	Row     int
	Col     int
}

func NewCharReader(input string) CharReader {
	return CharReader{content: input, Row: 0, Col: 0, i: -1}
}

func (r *CharReader) Peek() byte {
	return r.content[r.i+1]
}

func (r *CharReader) Next() byte {
	r.i++
	return r.content[r.i]
}

func (r *CharReader) IsEOF() bool {
	return r.i >= len(r.content)-1
}
