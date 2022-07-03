package com.github.panda_io.micro_panda.parser;

import java.util.List;

import com.github.panda_io.micro_panda.scanner.*;
import com.github.panda_io.micro_panda.ast.Program;

public class Parser {
    int position;
    Token token;
    String literal;

    Program program;
    //TO-DO scanner *scanner.Scanner

    public Parser(List<String> flags, Program program) {
        this.program = program;
        //scanner: scanner.NewScanner(flags), TO-DO
    }
    
    /*
    func (p *Parser) ParseBytes(source []byte) {
        file := token.NewFile("<input>"+fmt.Sprintf("%x", md5.Sum(source)), len(source))
        p.setSource(file, source)
        p.parseSourceFile(file)
    }

    func (p *Parser) ParseFile(file *token.File, source []byte) {
        p.setSource(file, source)
        p.parseSourceFile(file)
    }

    func (p *Parser) ParseExpression(source []byte) core.Expression {
        file := token.NewFile("<input>"+fmt.Sprintf("%x", md5.Sum(source)), len(source))
        p.setSource(file, source)
        return p.parseExpression()
    }

    func (p *Parser) ParseStatements(source []byte) core.Statement {
        file := token.NewFile("<input>"+fmt.Sprintf("%x", md5.Sum(source)), len(source))
        p.setSource(file, source)
        return p.parseBlockStatement()
    }

    func (p *Parser) next() {
        p.position, p.token, p.literal = p.scanner.Scan()
    }

    func (p *Parser) expect(t token.Token) {
        if p.token != t {
            p.expectedError(p.position, fmt.Sprintf("'%s'", t.String()))
        }
        p.next()
    }

    func (p *Parser) expectedError(position int, expect string) {
        expect = "expected " + expect
        if position == p.position {
            switch {
            case p.token == token.Semi && p.literal == "\n":
                expect += ", but found newline"
            case p.token.IsLiteral():
                expect += ", but found " + p.literal
            default:
                expect += ", but found '" + p.token.String() + "'"
            }
        }
        p.error(position, expect)
    }

    func (p *Parser) setSource(file *token.File, source []byte) {
        p.scanner.SetFile(file, source)
        p.next()
    }

    func (p *Parser) error(position int, message string) {
        panic(fmt.Sprintf("error: %s \n %s \n", p.scanner.Position(position).String(), message))
    }
    */
}
