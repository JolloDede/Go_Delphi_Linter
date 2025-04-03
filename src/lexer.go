package src

import (
	"slices"
	"strings"
)

type Lexer struct {
	reader CharReader
}

func NewLexer(input string) Lexer {
	return Lexer{reader: NewCharReader(input)}
}

func (l *Lexer) Peek() *Token {
	return l.next_token()
}

func (l *Lexer) Next() *Token {
	return l.next_token()
}

var operators = []byte{'+', '-', '*', '/', '=', ';', ':'}

func (l *Lexer) next_token() *Token {
	if l.reader.IsEOF() {
		return NewToken(EOFToken, "", l.reader.Row, l.reader.Col)
	}
	char := l.reader.Peek()

	if char == '\'' {
		return l.processStringLiteral()
	} else if '0' <= char && char <= '9' {
		return l.processNumber()
	} else if ('A' <= char && char <= 'Z') || ('a' <= char && char <= 'z') {
		// Parse code
		return l.processAlphanumeric()
		//
	} else if (char == '/' && l.reader.Peek_n(1) == '/') || char == '{' {
		char = l.reader.Peek_n(1)

		if char == '$' {
			return l.processConditionalComilations()
		} else {
			return l.processComment()
		}
	} else if slices.Contains(operators, char) {
		// process operators
		return l.processOperator()
	} else {
		panic("cannot parse, \"next_token\"")
	}

	return nil
}

func (l *Lexer) processComment() *Token {
	var sb strings.Builder
	var c byte

	c = l.reader.Peek()

	if c == '{' {
		for {
			c = l.reader.Next()

			if c == '}' {
				break
			}

			sb.WriteByte(c)
		}
	} else {
		c = l.reader.Peek_n(1)

		if c == '/' {
			l.reader.Jump(1)
			for {
				c = l.reader.Next()

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
	i := 1

	// Check multiline string
	for {
		c = l.reader.Peek_n(i)

		if c != '\'' {
			break
		}

		i++
	}

	if i > 1 {
		l.reader.Jump(i)
	}

	for {
		c = l.reader.Next()

		if c == '\'' {
			j := 1

			for {
				if l.reader.Peek_n(j) != '\'' {
					break
				}

				j++
			}

			if j%2 != 0 {
				if j != i {
					panic("Wrong string")
				}

				break
			} else {
				l.reader.Next()
			}
		}

		sb.WriteByte(c)
	}

	return NewToken(StringToken, sb.String(), l.reader.Row, l.reader.Col)
}

func (l *Lexer) processConditionalComilations() *Token {
	var sb strings.Builder

	for {
		c := l.reader.Next()

		if c == '}' {
			break
		}
		sb.WriteByte(c)
	}

	return NewToken(ConditionalCompilingToken, sb.String(), l.reader.Row, l.reader.Col)
}

func (l *Lexer) processNumber() *Token {
	var sb strings.Builder
	c := l.reader.Peek()

	sb.WriteByte(c)

	for {
		c = l.reader.Next()

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

func (l *Lexer) processOperator() *Token {
	c := l.reader.Peek()

	l.reader.Jump(1)

	return NewToken(OperatorToken, string(c), l.reader.Row, l.reader.Col)
}

var keywords = []string{"begin", "end", "unit", "interface", "implementation", "class", "record", "initialization", "finalization", "if", "then", "for, while", "do"}

func (l *Lexer) processAlphanumeric() *Token {
	var sb strings.Builder

	sb.WriteByte(l.reader.Peek())

	for {
		c := l.reader.Next()

		if slices.Contains([]byte{}, c) {

		}
		// maybe add some more this that can be in a identifier
		if ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') || (c == '_') {
			sb.WriteByte(c)
		} else {
			break
		}
	}

	if slices.Contains(keywords, sb.String()) {
		return NewToken(KeywordToken, sb.String(), l.reader.Row, l.reader.Col)
	} else {
		return NewToken(IdentifierToken, sb.String(), l.reader.Row, l.reader.Col)
	}
}
