package lex

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
				return l.pos, EOF, "EOF"
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
			sym, val := l.lexSlash(string(r))
			return l.pos, sym, val
		case '%':
			return l.pos, MOD, string(r)
		case '=':
			t, s := l.lexEquals(r)
			return l.pos, t, s
		case '!':
			return l.pos, NOT, string(r)
		case '<':
			return l.pos, LT, string(r)
		case '>':
			return l.pos, GT, string(r)
		case '(':
			return l.pos, LPAREN, string(r)
		case ')':
			return l.pos, RPAREN, string(r)
		case ',':
			return l.pos, COMMA, string(r)
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
		case '@':
			startPos := l.pos
			l.backup()
			lit := l.lexCompilerInstruction()
			if lit == "@import" {
				return startPos, IMPORT, lit
			} else {
				return startPos, ILLEGAL, lit
			}
		case '"':
			startPos := l.pos
			lit := l.lexString()
			return startPos, STRINGLITERAL, lit
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
						return startPos, kwmap[keyword], lit
					}
				}
				// need to check if it's a type annotation
				for _, tannot := range types {
					if tannot == lit {
						return startPos, TYPEANNOT, lit
					}
				}
				// need to check if it's a boolean value
				if lit == "true" {
					return startPos, TRUE, lit
				} else if lit == "false" {
					return startPos, FALSE, lit
				}
				return startPos, IDENT, lit
			} else {
				return l.pos, ILLEGAL, string(r)
			}
		}
		return l.pos, ILLEGAL, string(r)
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
			lit = lit + string(r)
		}
	}
}

func (l *Lexer) lexCompilerInstruction() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.pos.col++
		if r != ' ' {
			lit = lit + string(r)
		} else {
			return lit
		}
	}
}

func (l *Lexer) lexEquals(r rune) (Token, string) {
	s := string(r)
	r, _, _ = l.reader.ReadRune()
	s = s + string(r)
	switch r {
	case '=':
		return EQUALS, s
	default:
		return ASSIGN, s
	}
}

func (l *Lexer) lexBang(r rune) (Token, string) {
	s := string(r)
	r, _, _ = l.reader.ReadRune()
	s = s + string(r)
	switch r {
	case '=':
		return NOTEQUALS, s
	default:
		return NOT, s
	}
}

func (l *Lexer) lexSlash(r string) (Token, string) {
	lit := r
	for {
		r, _, _ := l.reader.ReadRune()
		if r == '/' {
			l.pos.col++
			// the rest of the line is a comment and should be ignored
			r, _, _ = l.reader.ReadRune()
			for r != '\n' {
				r, _, _ = l.reader.ReadRune()
			}
			l.resetPosition()
			return NEWLINE, "\n"
		} else {
			l.backup()
			return DIV, lit
		}
	}
}
