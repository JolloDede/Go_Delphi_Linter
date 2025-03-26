package src

import (
	"fmt"
	"strings"
)

type Lexer struct {
	reader CharReader
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
		return NewToken(EOFToken, "", l.reader.Row, l.reader.Col)
	}
	char, err := l.reader.Peek()

	if err != nil {
		fmt.Println(err)
		panic("How did we get here")
	}

	if char == '\'' {
		return l.processStringLiteral()
	} else if '0' <= char && char <= '9' {
		return l.processNumber()
	} else if ('A' <= char && char <= 'Z') || ('a' <= char && char <= 'z') {
		// Parse code
	} else if char == '/' || char == '{' {
		char, err = l.reader.Peek_n(1)
		if err != nil {
			panic("\"next_token\" Comment")
		}

		if char == '$' {
			return l.processConditionalComilations()
		} else {
			return l.processComment()
		}
	} else {
		panic("cannot parse, \"next_token\"")
	}

	return nil
}

func (l *Lexer) processComment() *Token {
	var sb strings.Builder
	var c byte
	var err error

	c, err = l.reader.Peek()

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
		c, err = l.reader.Peek_n(1)

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

	return NewToken(CommentToken, sb.String(), l.reader.Row, l.reader.Col)
}

func (l *Lexer) processStringLiteral() *Token {
	var sb strings.Builder
	var c byte
	var err error
	i := 0

	// Check multiline string
	for {
		c, err = l.reader.Peek_n(i)

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

	return NewToken(StringToken, sb.String(), l.reader.Row, l.reader.Col)
}

func (l *Lexer) processConditionalComilations() *Token {
	var sb strings.Builder

	for {
		c, err := l.reader.Next()

		if err != nil {
			panic("\"processConditionalComilations\" EOF before close ConditionalCompilation")
		}

		if c == '}' {
			break
		}
		sb.WriteByte(c)
	}

	return NewToken(ConditionalCompilingToken, sb.String(), l.reader.Row, l.reader.Col)
}

func (l *Lexer) processNumber() *Token {
	var sb strings.Builder
	c, err := l.reader.Peek()

	if err != nil {
		fmt.Println(err)
		panic("How did we get here \"processNumber\"")
	}

	sb.WriteByte(c)

	for {
		c, err = l.reader.Next()

		if err != nil {
			panic("\"processNumber\" EOF before close Number")
		}

		if '0' <= c && c <= '9' {
			sb.WriteByte(c)
		} else if c == '.' {
			sb.WriteByte(c)
		} else {
			break
		}
	}

	return NewToken(NumberToken, sb.String(), l.reader.Row, l.reader.Col)
}
