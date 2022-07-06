package com.github.panda_io.micro_panda.parser;

public class ModuleParser {
    /*

func (p *Parser) parseSourceFile(file *token.File) {
	m := &ast.Module{
		File: file,
	}
	m.Attributes = p.parseAttributes()
	m.Namespace = p.parseNamespace()
	m.Imports = p.parseImports()

	for p.token != token.EOF {
		attr := p.parseAttributes()
		public := p.parseModifier()
		switch p.token {
		case token.Const, token.Var:
			v := p.parseVariable(public, attr)
			v.Qualified = m.Namespace + "." + v.Name.Name
			if p.program.Declarations[v.Qualified] != nil {
				p.error(v.Name.GetPosition(), fmt.Sprintf("variable %s redeclared", v.Name.Name))
			}
			m.Variables = append(m.Variables, v)
			err := p.program.AddDeclaration(v)
			if err != nil {
				p.error(v.Name.GetPosition(), err.Error())
			}

		case token.Function:
			f := p.parseFunction(public, attr)
			f.Qualified = m.Namespace + "." + f.Name.Name
			if p.program.Declarations[f.Qualified] != nil {
				p.error(f.Name.GetPosition(), fmt.Sprintf("function %s redeclared", f.Name.Name))
			}
			m.Functions = append(m.Functions, f)
			err := p.program.AddDeclaration(f)
			if err != nil {
				p.error(f.Name.GetPosition(), err.Error())
			}

		case token.Enum:
			e := p.parseEnum(public, attr)
			e.Qualified = m.Namespace + "." + e.Name.Name
			if p.program.Declarations[e.Qualified] != nil {
				p.error(e.Name.GetPosition(), fmt.Sprintf("enum %s redeclared", e.Name.Name))
			}
			m.Enums = append(m.Enums, e)
			err := p.program.AddDeclaration(e)
			if err != nil {
				p.error(e.Name.GetPosition(), err.Error())
			}

		case token.Struct:
			s := p.parseStruct(public, attr)
			s.Qualified = m.Namespace + "." + s.Name.Name
			if p.program.Declarations[s.Qualified] != nil {
				p.error(s.Name.GetPosition(), fmt.Sprintf("class %s redeclared", s.Name.Name))
			}
			m.Structs = append(m.Structs, s)
			err := p.program.AddDeclaration(s)
			if err != nil {
				p.error(s.Name.GetPosition(), err.Error())
			}

		default:
			p.expectedError(p.position, "declaration")
		}
	}

	p.program.Modules[file.Name] = m
}

func (p *Parser) parseNamespace() string {
	p.expect(token.Namespace)
	if p.token == token.Semi {
		p.next()
		return core.Global
	}
	namespace := p.parseQualified()
	p.expect(token.Semi)
	return namespace
}

func (p *Parser) parseImports() []*ast.Import {
	imports := []*ast.Import{}
	for p.token == token.Import {
		p.expect(token.Import)
		u := &ast.Import{}
		u.Namespace = p.parseQualified()
		p.expect(token.Semi)
		imports = append(imports, u)
	}
	return imports
}
*/

}
