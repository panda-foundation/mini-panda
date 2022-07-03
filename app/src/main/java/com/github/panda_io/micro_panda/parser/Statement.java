package com.github.panda_io.micro_panda.parser;

public class Statement {
    /*

func (p *Parser) parseStatement() ast.Statement {
	switch p.token {
	case token.Break:
		s := &statement.Break{}
		s.SetPosition(p.position)
		p.next()
		p.expect(token.Semi)
		return s

	case token.Continue:
		s := &statement.Continue{}
		s.SetPosition(p.position)
		p.next()
		p.expect(token.Semi)
		return s

	case token.Return:
		s := &statement.Return{}
		s.SetPosition(p.position)
		p.next()
		if p.token != token.Semi {
			s.Expression = p.parseExpression()
		}
		p.expect(token.Semi)
		return s

	case token.LeftBrace:
		return p.parseBlockStatement()

	case token.If:
		return p.parseIfStatement()

	case token.Switch:
		return p.parseSwitchStatement()

	case token.For:
		return p.parseForStatement()

	default:
		return p.parseSimpleStatement(true)
	}
}

func (p *Parser) parseSimpleStatement(consumeSemi bool) ast.Statement {
	switch p.token {
	case token.Semi:
		s := &statement.Empty{}
		s.SetPosition(p.position)
		if consumeSemi {
			p.expect(token.Semi)
		}
		return s

	case token.Var:
		return p.parseDeclarationStatement(consumeSemi)

	case token.IDENT, token.This,
		token.CHAR, token.INT, token.FLOAT, token.STRING, token.BOOL, token.NULL, token.Void,
		token.LeftParen, token.LeftBracket,
		token.Plus, token.Minus, token.Not, token.BitXor:
		position := p.position
		x := p.parseExpression()
		if consumeSemi {
			p.expect(token.Semi)
		}
		s := &statement.ExpressionStatement{}
		s.SetPosition(position)
		s.Expression = x
		return s

	default:
		p.expectedError(p.position, "statement")
		return nil
	}
}

func (p *Parser) parseDeclarationStatement(consumeSemi bool) *statement.DeclarationStatement {
	s := &statement.DeclarationStatement{}
	s.SetPosition(p.position)
	p.next()
	s.Name = p.parseIdentifier()
	if p.token != token.Assign && p.token != token.Semi && p.token != token.Colon {
		s.Type = p.parseType()
	}
	if p.token == token.Assign {
		p.next()
		s.Value = p.parseExpression()
	}
	if consumeSemi {
		p.expect(token.Semi)
	}
	return s
}

func (p *Parser) parseBlockStatement() *statement.Block {
	s := &statement.Block{}
	s.SetPosition(p.position)
	p.next()
	for p.token != token.RightBrace {
		s.Statements = append(s.Statements, p.parseStatement())
	}
	p.next()
	return s
}

func (p *Parser) parseIfStatement() *statement.If {
	s := &statement.If{}
	p.next()
	p.expect(token.LeftParen)
	first := p.parseSimpleStatement(false)
	if p.token == token.Semi {
		p.next()
		s.Initialization = first
		condition := p.parseSimpleStatement(false)
		if expr, ok := condition.(*statement.ExpressionStatement); ok {
			s.Condition = expr.Expression
		} else {
			p.error(condition.GetPosition(), "if condition must be an expression")
		}
	} else {
		if expr, ok := first.(*statement.ExpressionStatement); ok {
			s.Condition = expr.Expression
		} else {
			p.error(first.GetPosition(), "if condition must be an expression")
		}
	}
	p.expect(token.RightParen)
	s.Body = p.parseStatement()
	if p.token == token.Else {
		p.next()
		if p.token == token.If {
			s.Else = p.parseIfStatement()
		} else {
			s.Else = p.parseStatement()
		}
	}
	return s
}

func (p *Parser) parseSwitchStatement() *statement.Switch {
	s := &statement.Switch{}
	s.SetPosition(p.position)
	p.next()
	p.expect(token.LeftParen)
	first := p.parseSimpleStatement(false)
	var operand ast.Statement
	if p.token == token.Semi {
		p.next()
		s.Initialization = first
		operand = p.parseSimpleStatement(false)
	} else {
		operand = first
	}
	if expr, ok := operand.(*statement.ExpressionStatement); ok {
		s.Operand = expr.Expression
	} else {
		p.error(operand.GetPosition(), "expect expression")
	}

	p.expect(token.RightParen)
	p.expect(token.LeftBrace)
	for p.token == token.Case {
		s.Cases = append(s.Cases, p.parseCaseStatement())
	}
	if p.token == token.Default {
		s.Default = p.parseCaseStatement()
	}
	if len(s.Cases) == 0 {
		p.error(s.GetPosition(), "expect 'case'")
	}
	p.expect(token.RightBrace)
	return s
}

func (p *Parser) parseCaseStatement() *statement.Case {
	s := &statement.Case{}
	s.SetPosition(p.position)
	s.Token = p.token
	if p.token == token.Case {
		p.next()
		s.Case = p.parseExpression()
	} else {
		p.expect(token.Default)
	}
	p.expect(token.Colon)
	s.Body = p.parseStatement()
	return s
}

// for {}
// for (condition) {}
// for (init; condition; post) {}
func (p *Parser) parseForStatement() ast.Statement {
	position := p.position
	p.next()
	if p.token != token.LeftParen {
		s := &statement.For{}
		s.SetPosition(position)
		s.Body = p.parseStatement()
		return s
	} else {
		p.next()
		first := p.parseSimpleStatement(false)
		if p.token == token.RightParen {
			p.next()
			s := &statement.For{}
			s.SetPosition(position)
			if expr, ok := first.(*statement.ExpressionStatement); ok {
				s.Condition = expr.Expression
			} else {
				p.error(first.GetPosition(), "expect expression")
			}
			s.Body = p.parseStatement()
			return s
		} else {
			p.expect(token.Semi)
			s := &statement.For{}
			s.Initialization = first
			second := p.parseSimpleStatement(false)
			if expr, ok := second.(*statement.ExpressionStatement); ok {
				s.Condition = expr.Expression
			} else {
				p.error(second.GetPosition(), "expect expression")
			}
			p.expect(token.Semi)
			s.Post = p.parseSimpleStatement(false)
			p.expect(token.RightParen)
			s.Body = p.parseStatement()
			return s
		}
	}
}
*/
}
