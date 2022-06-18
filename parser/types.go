package parser

import (
	"strconv"

	"github.com/panda-io/micro-panda/ast/core"
	"github.com/panda-io/micro-panda/ast/expression"
	"github.com/panda-io/micro-panda/token"
)

func (p *Parser) parseType() core.Type {
	if p.token.IsScalar() {
		t := &core.TypeBuiltin{}
		t.Position = p.position
		t.Token = p.token
		p.next()
		return t
	}
	if p.token == token.Function {
		p.next()
		return p.parseFunctionType()
	}
	if p.token == token.LeftBracket {
		return p.parseTypeArray()
	}
	if p.token == token.Pointer {
		p.next()
		return p.parseTypePointer()
	}
	return p.parseTypeName()
}

func (p *Parser) parseTypeArray() *core.TypeArray {
	t := &core.TypeArray{}
	t.Position = p.position
	for p.token == token.LeftBracket {
		p.next()
		count := 0
		if p.token == token.INT {
			count, _ = strconv.Atoi(p.literal)
			if count < 1 {
				p.error(p.position, "array count must > 0")
			}
			p.next()
		}
		t.Dimension = append(t.Dimension, count)
		p.expect(token.RightBracket)
	}
	t.ElementType = p.parseType()
	return t
}

func (p *Parser) parseTypeName() *core.TypeName {
	t := &core.TypeName{}
	t.Position = p.position
	t.Name = p.parseIdentifier().Name
	if p.token == token.Dot {
		p.next()
		t.Selector = t.Name
		t.Name = p.parseIdentifier().Name
	}
	return t
}

func (p *Parser) parseTypePointer() *core.TypePointer {
	t := &core.TypePointer{}
	if p.token == token.Less {
		p.next()
		t.ElementType = p.parseType()
		p.expect(token.Greater)
	} else {
		t.ElementType = core.TypeU8
	}
	return t
}

func (p *Parser) parseParameters() []*expression.Parameter {
	t := []*expression.Parameter{}
	p.expect(token.LeftParen)
	if p.token == token.RightParen {
		p.next()
		return nil
	}
	t = append(t, p.parseParameter())
	for p.token == token.Comma {
		p.next()
		t = append(t, p.parseParameter())
	}
	p.expect(token.RightParen)
	return t
}

func (p *Parser) parseParameter() *expression.Parameter {
	t := &expression.Parameter{}
	t.Position = p.position
	t.Name = p.parseIdentifier().Name
	t.Type = p.parseType()
	return t
}

func (p *Parser) parseArguments() *expression.Arguments {
	t := &expression.Arguments{}
	t.Position = p.position
	p.expect(token.LeftParen)
	if p.token == token.RightParen {
		p.next()
		return t
	}
	t.Arguments = append(t.Arguments, p.parseExpression())
	for p.token == token.Comma {
		p.next()
		t.Arguments = append(t.Arguments, p.parseExpression())
	}
	p.expect(token.RightParen)
	return t
}

func (p *Parser) parseFunctionType() *core.TypeFunction {
	t := &core.TypeFunction{}
	t.Position = p.position
	p.expect(token.LeftParen)
	if p.token == token.RightParen {
		p.next()
		return t
	}
	t.Parameters = append(t.Parameters, p.parseType())
	for p.token == token.Comma {
		p.next()
		t.Parameters = append(t.Parameters, p.parseType())
	}
	p.expect(token.RightParen)
	if p.token != token.Semi && p.token != token.Assign {
		t.ReturnType = p.parseType()
	}
	return t
}
