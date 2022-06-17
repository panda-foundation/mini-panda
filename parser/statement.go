package parser

import (
	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.token {
	case token.Break:
		s := &ast.Break{}
		s.Position = p.position
		p.next()
		p.expect(token.Semi)
		return s

	case token.Continue:
		s := &ast.Continue{}
		s.Position = p.position
		p.next()
		p.expect(token.Semi)
		return s

	case token.Return:
		s := &ast.Return{}
		s.Position = p.position
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
		s := &ast.Empty{}
		s.Position = p.position
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
		s := &ast.ExpressionStatement{}
		s.Position = position
		s.Expression = x
		return s

	default:
		p.expectedError(p.position, "statement")
		return nil
	}
}

func (p *Parser) parseDeclarationStatement(consumeSemi bool) *ast.DeclarationStatement {
	s := &ast.DeclarationStatement{}
	s.Position = p.position
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

func (p *Parser) parseBlockStatement() *ast.Block {
	s := &ast.Block{}
	s.Position = p.position
	p.next()
	for p.token != token.RightBrace {
		s.Statements = append(s.Statements, p.parseStatement())
	}
	p.next()
	return s
}

func (p *Parser) parseIfStatement() *ast.If {
	s := &ast.If{}
	p.next()
	p.expect(token.LeftParen)
	first := p.parseSimpleStatement(false)
	if p.token == token.Semi {
		p.next()
		s.Initialization = first
		condition := p.parseSimpleStatement(false)
		if expr, ok := condition.(*ast.ExpressionStatement); ok {
			s.Condition = expr.Expression
		} else {
			p.error(condition.GetPosition(), "if condition must be an expression")
		}
	} else {
		if expr, ok := first.(*ast.ExpressionStatement); ok {
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

func (p *Parser) parseSwitchStatement() *ast.Switch {
	s := &ast.Switch{}
	s.Position = p.position
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
	if expr, ok := operand.(*ast.ExpressionStatement); ok {
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
		p.error(s.Position, "expect 'case'")
	}
	p.expect(token.RightBrace)
	return s
}

func (p *Parser) parseCaseStatement() *ast.Case {
	s := &ast.Case{}
	s.Position = p.position
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
		s := &ast.For{}
		s.Position = position
		s.Body = p.parseStatement()
		return s
	} else {
		p.next()
		first := p.parseSimpleStatement(false)
		if p.token == token.RightParen {
			p.next()
			s := &ast.For{}
			s.Position = position
			if expr, ok := first.(*ast.ExpressionStatement); ok {
				s.Condition = expr.Expression
			} else {
				p.error(first.GetPosition(), "expect expression")
			}
			s.Body = p.parseStatement()
			return s
		} else {
			p.expect(token.Semi)
			s := &ast.For{}
			s.Initialization = first
			second := p.parseSimpleStatement(false)
			if expr, ok := second.(*ast.ExpressionStatement); ok {
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
