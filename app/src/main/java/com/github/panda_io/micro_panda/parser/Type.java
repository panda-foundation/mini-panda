package com.github.panda_io.micro_panda.parser;

public class Type {
    /*

func (p *Parser) parseType() ast.Type {
	if p.token.IsScalar() {
		t := &ast_types.TypeBuiltin{}
		t.SetPosition(p.position)
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

func (p *Parser) parseTypeArray() *ast_types.TypeArray {
	t := &ast_types.TypeArray{}
	t.SetPosition(p.position)
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

func (p *Parser) parseTypeName() *ast_types.TypeName {
	t := &ast_types.TypeName{}
	t.SetPosition(p.position)
	qualified := p.parseQualified()
	if strings.Contains(qualified, ".") {
		t.Qualified = qualified
		names := strings.Split(qualified, ".")
		t.Name = names[len(names)-1]
	} else {
		t.Name = qualified
	}
	return t
}

func (p *Parser) parseTypePointer() *ast_types.TypePointer {
	t := &ast_types.TypePointer{}
	if p.token == token.Less {
		p.next()
		t.ElementType = p.parseType()
		p.expect(token.Greater)
	} else {
		t.ElementType = ast_types.TypeU8
	}
	return t
}

func (p *Parser) parseParameters() []*declaration.Parameter {
	t := []*declaration.Parameter{}
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

func (p *Parser) parseParameter() *declaration.Parameter {
	t := &declaration.Parameter{}
	t.SetPosition(p.position)
	t.Name = p.parseIdentifier().Name
	t.Typ = p.parseType()
	return t
}

func (p *Parser) parseArguments() []ast.Expression {
	expressions := []ast.Expression{}
	p.expect(token.LeftParen)
	if p.token == token.RightParen {
		p.next()
		return expressions
	}
	expressions = append(expressions, p.parseExpression())
	for p.token == token.Comma {
		p.next()
		expressions = append(expressions, p.parseExpression())
	}
	p.expect(token.RightParen)
	return expressions
}

func (p *Parser) parseFunctionType() *ast_types.TypeFunction {
	t := &ast_types.TypeFunction{}
	t.SetPosition(p.position)
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
*/
}
