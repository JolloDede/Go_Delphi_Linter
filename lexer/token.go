package lexer

type TokenTyp string

const (
	CommentToken              TokenTyp = "Comment"
	StringToken               TokenTyp = "String"
	NumberToken               TokenTyp = "Number"
	ConditionalCompilingToken TokenTyp = "CondComp"
	EOFToken                  TokenTyp = "EOF"
	OperatorToken             TokenTyp = "Token"
	KeywordToken              TokenTyp = "Keyword"
	IdentifierToken           TokenTyp = "Identifier"
	WhiteSpaceToken           TokenTyp = "Whitespace"
)

type Token struct {
	Typ     TokenTyp
	Content string

	row int
	col int
}

func NewToken(typ TokenTyp, content string, row int, col int) *Token {
	return &Token{Typ: typ, Content: content, row: row, col: col}
}
