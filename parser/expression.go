package parser

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/token"
)

func (p *Parser) parseExpression() ast.Expression {
	return p.parseBinaryExpression(0)
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	e := &ast.Identifier{}
	e.Position = p.position
	if p.token == token.IDENT {
		e.Name = p.literal
		p.next()
	} else {
		p.expect(token.IDENT)
	}
	return e
}

func (p *Parser) parseOperand() ast.Expression {
	switch p.token {
	case token.IDENT:
		return p.parseIdentifier()

	case token.Bool, token.Int8, token.Int16, token.Int32, token.Int64,
		token.Uint8, token.Uint16, token.Uint32, token.Uint64,
		token.Float16, token.Float32, token.Float64, token.Pointer, token.LeftBracket:
		e := &ast.Conversion{}
		e.Position = p.position
		e.Typ = p.parseType()
		p.expect(token.LeftParen)
		e.Value = p.parseExpression()
		p.expect(token.RightParen)
		return e

	case token.CHAR, token.INT, token.FLOAT, token.STRING, token.BOOL, token.NULL, token.Void:
		e := &ast.Literal{}
		e.Position = p.position
		e.Token = p.token
		e.Value = p.literal
		p.next()
		return e

	case token.LeftParen:
		e := &ast.Parentheses{}
		e.Position = p.position
		p.next()
		e.Expression = p.parseExpression()
		p.expect(token.RightParen)
		return e

	case token.LeftBrace:
		e := &ast.Initializer{}
		e.Position = p.position
		p.next()
		for {
			e.Expressions = append(e.Expressions, p.parseExpression())
			if p.token == token.Comma {
				p.next()
			} else if p.token == token.RightBrace {
				p.next()
				break
			} else {
				p.error(p.position, "unexpected "+p.token.String())
				return nil
			}
		}
		return e

	case token.This:
		e := &ast.This{}
		e.Position = p.position
		p.next()
		return e

	case token.Sizeof:
		e := &ast.Sizeof{}
		e.Position = p.position
		p.next()
		p.expect(token.LeftParen)
		e.Target = p.parseType()
		p.expect(token.RightParen)
		return e

	default:
		p.error(p.position, "unexpected "+p.token.String())
		return nil
	}
}

func (p *Parser) parsePrimaryExpression() ast.Expression {
	x := p.parseOperand()
	for {
		switch p.token {
		case token.Dot:
			e := &ast.MemberAccess{}
			e.Position = p.position
			p.next()
			e.Parent = x
			e.Member = p.parseIdentifier()
			x = e

		case token.LeftBracket:
			e := &ast.Subscripting{}
			e.Position = p.position
			e.Parent = x
			for p.token == token.LeftBracket {
				p.next()
				e.Indexes = append(e.Indexes, p.parseExpression())
				p.expect(token.RightBracket)
			}
			x = e

		case token.LeftParen:
			e := &ast.Invocation{}
			e.Position = p.position
			e.Function = x
			e.Arguments = p.parseArguments()
			x = e

		case token.PlusPlus:
			e := &ast.Increment{}
			e.Position = p.position
			e.Expression = x
			p.next()
			return e

		case token.MinusMinus:
			e := &ast.Decrement{}
			e.Position = p.position
			e.Expression = x
			p.next()
			return e

		default:
			return x
		}
	}
}

func (p *Parser) parseUnaryExpression() ast.Expression {
	switch p.token {
	case token.Plus, token.Minus, token.Not, token.Complement, token.BitAnd:
		e := &ast.Unary{}
		e.Position = p.position
		e.Operator = p.token
		p.next()
		e.Expression = p.parseUnaryExpression()
		return e

	default:
		return p.parsePrimaryExpression()
	}
}

func (p *Parser) parseBinaryExpression(precedence int) ast.Expression {
	x := p.parseUnaryExpression()
	for {
		if p.token == token.Semi {
			return x
		}
		op := p.token
		opPrec := p.token.Precedence()
		if opPrec <= precedence {
			return x
		}
		p.next()
		y := p.parseBinaryExpression(opPrec)
		x = &ast.Binary{
			Left:     x,
			Operator: op,
			Right:    y,
		}
	}
}
