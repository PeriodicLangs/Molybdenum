package ast

type Program struct {
	Statements []Statement
}

func (p *Program) Literal() string {
	if len(p.Statements) > 0 {
		s := ""
		for _, stmt := range p.Statements {
			s += "\n" + stmt.Literal()
		}
		return s
	} else {
		return ""
	}
}
