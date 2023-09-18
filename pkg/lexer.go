package pkg

type Lexer struct {
	input        string
	position     int
	nextPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = Token{Type: Equal, Literal: literal}
		} else {
			tok = newToken(Assign, l.ch)
		}
	case '+':
		tok = newToken(Plus, l.ch)
	case '-':
		tok = newToken(Minus, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = Token{Type: NotEqual, Literal: literal}
		} else {
			tok = newToken(Negation, l.ch)
		}
	case '*':
		tok = newToken(Multiplication, l.ch)
	case '/':
		tok = newToken(Division, l.ch)
	case '<':
		tok = newToken(LessThan, l.ch)
	case '>':
		tok = newToken(GreaterThan, l.ch)
	case ',':
		tok = newToken(Comma, l.ch)
	case ';':
		tok = newToken(Semicolon, l.ch)
	case ':':
		tok = newToken(Colon, l.ch)
	case '(':
		tok = newToken(LeftParen, l.ch)
	case ')':
		tok = newToken(RightParen, l.ch)
	case '{':
		tok = newToken(LeftBrace, l.ch)
	case '}':
		tok = newToken(RightBrace, l.ch)
	case '[':
		tok = newToken(LeftBracket, l.ch)
	case ']':
		tok = newToken(RightBracket, l.ch)
	case '"':
		tok.Type = String
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdentifierType(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = Int
			return tok
		} else {
			tok = newToken(Illegal, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readChar() {
	l.ch = l.peekChar()
	l.position = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) readString() string {
	pos := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}
