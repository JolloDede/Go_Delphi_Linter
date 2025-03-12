package src

import (
	"fmt"
	"strings"
)

type Lexer struct {
	reader CharReader
}

type TokenTyp string

const (
	CommentToken TokenTyp = "Comment"
	StringToken  TokenTyp = "String"
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
	char, err := l.reader.Peek(0)

	if err != nil {
		fmt.Println(err)
		panic("How did we get here")
	}

	if char == '\'' {
		// Parse string
		return l.processStringLiteral()
	} else if '0' <= char && char <= '9' {
		// Parse numeric
	} else if ('A' <= char && char <= 'Z') || ('a' <= char && char <= 'z') {
		// Parse code
	} else if char == '/' || char == '{' {
		// Parse comment
		return l.processComment()
	} else {
		panic("cannot parse, \"next_token\"")
	}

	return nil
}

func (l *Lexer) processComment() *Token {
	var sb strings.Builder
	var c byte
	var err error

	c, err = l.reader.Peek(0)

	if err != nil {
		fmt.Println(err)
		panic("How did we get here \"processComment\"")
	}

	if c == '{' {
		for {
			c, err = l.reader.Next()

			if err != nil {
				fmt.Println(err)
				panic("How did we get here \"processComment\"")
			}

			if c == '}' {
				break
			}

			sb.WriteByte(c)
		}
	} else {
		c, err = l.reader.Peek(1)

		if err != nil {
			fmt.Println(err)
			panic("How did we get here \"processComment\"")
		}

		if c == '/' {
			l.reader.Jump(1)
			for {
				c, err = l.reader.Next()

				if err != nil {
					fmt.Println(err)
					panic("How did we get here \"processComment\"")
				}

				if c == '\n' {
					break
				}

				sb.WriteByte(c)
			}
		}
	}

	return &Token{Typ: CommentToken, content: sb.String()}
}

func (l *Lexer) processStringLiteral() *Token {
	var sb strings.Builder
	var c byte
	var err error
	i := 0

	// Check multiline string
	for {
		c, err = l.reader.Peek(i)

		if err != nil {
			fmt.Println(err)
			panic("How did we get here \"processStringLiteral\"")
		}

		if c != '\'' && i > 1 && i%2 != 0 {
			break
		}

		i++
	}

	l.reader.Jump(i)

	for {
		c, err = l.reader.Next()

		if err != nil {
			fmt.Println(err)
			panic("How did we get here \"processStringLiteral\"")
		}

		if c == '\'' {
			break
		}

		sb.WriteByte(c)
	}

	return &Token{Typ: StringToken, content: sb.String()}
}
