package com.github.panda_io.micro_panda.parser;

public class Expression {
    /*

func (p *Parser) parseExpression() ast.Expression {
	return p.parseBinaryExpression(0)
}

func (p *Parser) parseIdentifier() *expression.Identifier {
	e := &expression.Identifier{}
	e.SetPosition(p.position)
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
		e := &expression.Conversion{}
		e.SetPosition(p.position)
		e.Typ = p.parseType()
		p.expect(token.LeftParen)
		e.Value = p.parseExpression()
		p.expect(token.RightParen)
		return e

	case token.CHAR, token.INT, token.FLOAT, token.STRING, token.BOOL, token.NULL, token.Void:
		e := &expression.Literal{}
		e.SetPosition(p.position)
		e.Token = p.token
		e.Value = p.literal
		p.next()
		return e

	case token.LeftParen:
		e := &expression.Parentheses{}
		e.SetPosition(p.position)
		p.next()
		e.Expression = p.parseExpression()
		p.expect(token.RightParen)
		return e

	case token.LeftBrace:
		e := &expression.Initializer{}
		e.SetPosition(p.position)
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
		e := &expression.This{}
		e.SetPosition(p.position)
		p.next()
		return e

	case token.Sizeof:
		e := &expression.Sizeof{}
		e.SetPosition(p.position)
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
			e := &expression.MemberAccess{}
			e.SetPosition(p.position)
			p.next()
			e.Parent = x
			e.Member = p.parseIdentifier()
			x = e

		case token.LeftBracket:
			e := &expression.Subscripting{}
			e.SetPosition(p.position)
			e.Parent = x
			for p.token == token.LeftBracket {
				p.next()
				e.Indexes = append(e.Indexes, p.parseExpression())
				p.expect(token.RightBracket)
			}
			x = e

		case token.LeftParen:
			e := &expression.Invocation{}
			e.SetPosition(p.position)
			e.Function = x
			e.Arguments = p.parseArguments()
			x = e

		case token.PlusPlus:
			e := &expression.Increment{}
			e.SetPosition(p.position)
			e.Expression = x
			p.next()
			return e

		case token.MinusMinus:
			e := &expression.Decrement{}
			e.SetPosition(p.position)
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
	case token.Plus, token.Minus, token.Not, token.Complement, token.BitAnd, token.Mul:
		e := &expression.Unary{}
		e.SetPosition(p.position)
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
		x = &expression.Binary{
			Left:     x,
			Operator: op,
			Right:    y,
		}
	}
}
*/
}
