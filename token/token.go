package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	Illegal = "Illegal"
	EOF     = "EOF"

	// Identifiers + Literals
	Identifier = "Identifier" // add, x ,y, ...
	Int        = "Int"        // 123456
	String     = "String"     // "x", "y"

	// Operadores Matematicos
	Plus           = "+"
	Minus          = "-"
	Multiplication = "*"
	Division       = "/"

	// Operadores

	Assign   = "="
	Negation = "!"
	Equal    = "=="
	NotEqual = "!="

	LessThan    = "<"
	GreaterThan = ">"

	// Delimitadores
	Comma     = ","
	Semicolon = ";"
	Colon     = ":"

	LeftParen    = "("
	RightParen   = ")"
	LeftBrace    = "{"
	RightBrace   = "}"
	LeftBracket  = "["
	RightBracket = "]"

	// Keywords
	Function = "Function"
	Let      = "Let"
	True     = "True"
	False    = "False"
	If       = "If"
	Else     = "Else"
	Return   = "Return"
)

var keywords = map[string]TokenType{
	"fn":      Function,
	"var":     Let,
	"verdd":   True,
	"falso":   False,
	"si":      If,
	"si-no":   Else,
	"regresa": Return,
}

func LookupIdentifierType(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return Identifier
}
