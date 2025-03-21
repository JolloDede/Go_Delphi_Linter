package src

type TokenTyp string

const (
	CommentToken              TokenTyp = "Comment"
	StringToken               TokenTyp = "String"
	ConditionalCompilingToken TokenTyp = "CondComp"
)

type Token struct {
	Typ     TokenTyp
	content string

	row int
	col int
}

func NewToken(typ TokenTyp, content string, row int, col int) *Token {
	return &Token{Typ: typ, content: content, row: row, col: col}
}
