package parser

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/token"
)

func (p *Parser) parseVariable(public bool, attributes []*ast.Attribute) *ast.Variable {
	v := &ast.Variable{}
	v.Public = public
	v.Attributes = attributes
	v.Token = p.token
	if p.token == token.Const {
		v.Const = true
	}
	p.next()
	v.Name = p.parseIdentifier()
	v.Type = p.parseType()
	if p.token == token.Assign {
		p.next()
		v.Value = p.parseExpression()
	}
	if v.Const && v.Value == nil {
		p.error(v.Name.Position, "constant declaration must be initalized")
	}
	p.expect(token.Semi)
	return v
}

func (p *Parser) parseFunction(public bool, attributes []*ast.Attribute) *ast.Function {
	d := &ast.Function{
		Type: &ast.TypeFunction{},
	}
	d.Public = public
	d.Attributes = attributes
	p.next()
	d.Name = p.parseIdentifier()
	d.Parameters = p.parseParameters()
	if p.token != token.Semi && p.token != token.LeftBrace {
		d.ReturnType = p.parseType()
	}
	if p.token == token.LeftBrace {
		d.Body = p.parseBlockStatement()
	} else if p.token == token.Semi {
		p.next()
	}
	return d
}

func (p *Parser) parseEnum(public bool, attributes []*ast.Attribute) *ast.Enum {
	e := &ast.Enum{}
	e.Public = public
	e.Attributes = attributes
	p.next()
	e.Name = p.parseIdentifier()
	p.expect(token.LeftBrace)
	for p.token != token.RightBrace {
		v := &ast.Variable{}
		v.Const = true
		v.Name = p.parseIdentifier()
		if p.token == token.Assign {
			p.next()
			v.Value = p.parseExpression()
		}
		err := e.AddMember(v)
		if err != nil {
			p.error(v.Position, err.Error())
		}
		if p.token != token.Comma {
			break
		}
		p.next()
	}
	p.expect(token.RightBrace)
	return e
}

func (p *Parser) parseStruct(public bool, attributes []*ast.Attribute) *ast.Struct {
	s := &ast.Struct{}
	s.Public = public
	s.Attributes = attributes
	p.next()
	s.Name = p.parseIdentifier()

	p.expect(token.LeftBrace)
	for p.token != token.RightBrace {
		attr := p.parseAttributes()
		modifier := p.parseModifier()
		switch p.token {
		case token.Const, token.Var:
			v := p.parseVariable(modifier, attr)
			err := s.AddVariable(v)
			if err != nil {
				p.error(v.Position, err.Error())
			}

		case token.Function:
			f := p.parseFunction(modifier, attr)
			err := s.AddFunction(f)
			if err != nil {
				p.error(f.Position, err.Error())
			}

		default:
			p.expectedError(p.position, "member declaration")
		}
	}
	p.expect(token.RightBrace)
	return s
}

func (p *Parser) parseModifier() bool {
	if p.token == token.Public {
		p.next()
		return true
	}
	return false
}

func (p *Parser) parseAttributes() []*ast.Attribute {
	if p.token != token.META {
		return nil
	}
	var attr []*ast.Attribute
	for p.token == token.META {
		p.next()
		if p.token != token.IDENT {
			p.expect(token.IDENT)
		}
		m := &ast.Attribute{Position: p.position}
		m.Name = p.literal
		p.next()

		if p.token == token.STRING {
			m.Text = p.literal
			p.next()
		} else if p.token == token.LeftParen {
			p.next()
			if p.token == token.STRING {
				m.Text = p.literal
				p.next()
			} else {
				m.Values = make(map[string]*ast.Literal)
				for {
					if p.token == token.IDENT {
						name := p.literal
						p.next()
						p.expect(token.Assign)
						switch p.token {
						case token.CHAR, token.INT, token.FLOAT, token.STRING, token.BOOL:
							if _, ok := m.Values[name]; ok {
								p.error(p.position, "duplicated attribute "+name)
							}
							m.Values[name] = &ast.Literal{
								Token: p.token,
								Value: p.literal,
							}
							m.Values[name].Position = p.position
						default:
							p.expectedError(p.position, "basic literal (bool, char, int, float, string)")
						}
						p.next()
						if p.token == token.RightParen {
							break
						}
						p.expect(token.Comma)
					} else {
						p.expect(token.IDENT)
					}
				}
			}
			p.expect(token.RightParen)
		}
		attr = append(attr, m)
	}
	return attr
}
