package main

import (
	"bufio"
	"io"
	"unicode"
)

type Position struct {
	line int
	col  int
}

type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		reader: bufio.NewReader(reader),
		pos:    Position{line: 1, col: 0},
	}
}

func (l *Lexer) Lex() (Position, Token, string) {
	// keep looping until we return a token
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}

			// at this point there isn't much we can do, and the compiler
			// should just return the raw error to the user
			panic(err)
		}

		l.pos.col++

		switch r {
		case '\n':
			l.resetPosition()
			return l.pos, NEWLINE, string(r)
		case '+':
			return l.pos, ADD, string(r)
		case '*':
			return l.pos, MUL, string(r)
		case '-':
			return l.pos, SUB, string(r)
		case '/':
			return l.pos, DIV, string(r)
		case '%':
			return l.pos, MOD, string(r)
		case '=':
			return l.pos, ASSIGN, string(r)
		case '(':
			return l.pos, LPAREN, string(r)
		case ')':
			return l.pos, RPAREN, string(r)
		case '[':
			return l.pos, LSQRBRAC, string(r)
		case ']':
			return l.pos, RSQRBRAC, string(r)
		case '.':
			return l.pos, DOT, string(r)
		case '{':
			return l.pos, BLOCKSTART, string(r)
		case '}':
			return l.pos, BLOCKEND, string(r)
		default:
			if unicode.IsSpace(r) {
				continue
			} else if unicode.IsDigit(r) {
				startPos := l.pos
				l.backup()
				lit := l.lexInt()
				return startPos, INTLITERAL, lit
			} else if unicode.IsLetter(r) {
				startPos := l.pos
				l.backup()
				lit := l.lexIdent()
				// need to check if it's a keyword
				for _, keyword := range keywords {
					if keyword == lit {
						return startPos, KEYWORD, lit
					}
				}
				return startPos, IDENT, lit
			} else if r == '"' {
				startPos := l.pos
				l.backup()
				lit := l.lexString()
				return startPos, STRINGLITERAL, lit
			} else {
				return l.pos, ILLEGAL, string(r)
			}
		}
	}
}

func (l *Lexer) resetPosition() {
	l.pos.line++
	l.pos.col = 0
}

func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}

	l.pos.col--
}

func (l *Lexer) lexInt() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.pos.col++
		if unicode.IsDigit(r) {
			lit = lit + string(r)
		} else {
			l.backup()
			return lit
		}
	}
}

func (l *Lexer) lexIdent() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.pos.col++
		if unicode.IsLetter(r) {
			lit = lit + string(r)
		} else {
			l.backup()
			return lit
		}
	}
}

// not working! fix!
func (l *Lexer) lexString() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.pos.col++
		if r == '"' {
			return lit
		} else {
			l.backup()
			lit = lit + string(r)
		}
	}
}
