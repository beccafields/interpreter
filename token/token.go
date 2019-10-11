package token

// TokenType is the type of the token
type TokenType string

// Token struct hold the type and value
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifier + literals
	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"   // 134356

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	EQ       = "=="
	NOT_EQ   = "!="

	LT = "<"
	GT = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent checks the keywords table to see if the given identifier is a keyword.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		// return the keyword's TokenType constant
		return tok
	}
	// user definied identifier
	return IDENT
}
