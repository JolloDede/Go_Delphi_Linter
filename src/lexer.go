package src

import "strings"

type Lexer struct {
	reader CharReader
}

type TokenTyp string

const (
	CommentToken TokenTyp = "Comment"
)

type Token struct {
	Typ     TokenTyp
	content string

	row int
	col int
}

func NewLexer(input string) Lexer {
	return Lexer{reader: NewCharReader(input)}
}

func (l *Lexer) Peek() *Token {
	// return r.content[r.i+1]
	return l.next_token()
}

func (l *Lexer) Next() *Token {
	// 	r.i++
	// 	return r.content[r.i]
	return l.next_token()
}

func (l *Lexer) next_token() *Token {
	if l.reader.IsEOF() {
		return &Token{Typ: "EOF"}
	}
	char := l.reader.Peek()

	if '0' <= char && char <= '9' {
		// Parse numeric
	} else if ('A' <= char && char <= 'Z') || ('a' <= char && char <= 'z') {
		// Parse code
	} else if char == '/' {
		// Parse comment
		// return processComment()
	} else {
		panic("cannot parse")
	}

	return nil
}

func (l *Lexer) processComment() Token {
	var sb strings.Builder

	sb.WriteByte(l.reader.Peek())
	for l.reader.Next() != '\n' {
		l.reader.Next()
	}

	return Token{Typ: CommentToken, content: sb.String()}
}
