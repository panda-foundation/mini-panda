package com.github.panda_io.micro_panda.parser;

import com.github.panda_io.micro_panda.ast.Program;
import com.github.panda_io.micro_panda.ast.expression.Identifier;
import com.github.panda_io.micro_panda.scanner.Scanner;
import com.github.panda_io.micro_panda.scanner.Token;

public class Context {
    Program program;
    Scanner scanner;

    int position;
	Token token;
	String literal;

    public Context(Program program, Scanner scanner) {
        this.program = program;
        this.scanner = scanner;
    }

    public void next() throws Exception {
        Scanner.Result result = this.scanner.scan();
        this.position = result.position;
        this.token = result.token;
        this.literal = result.literal;
    }

    void expect(Token token) throws Exception {
        if (this.token != token) {
            this.unexpected(this.position, String.format("'%s'", token.toString()));
        }
        this.next();
    }

    void unexpected(int position, String expect) throws Exception {
        expect = "expected " + expect;
        if (position == this.position) {
            if (this.token == Token.Semi && this.literal == "\n") {
                expect += ", but found newline";
            } else if (this.token.isLiteral()){
                expect += ", but found " + this.literal;
            } else {
                expect += ", but found '" + this.token.toString() + "'";
            }
        }
        this.program.printLocation(position);
        throw new RuntimeException(expect);
    }

	void addError(int offset, String message) {
        this.program.addError(offset, message);
    }

    Identifier  parseIdentifier() throws Exception {
        Identifier identifier = new Identifier();
        identifier.setOffset(this.position);
        if (this.token == Token.IDENT) {
            identifier.name = this.literal;
            this.next();
        } else {
            this.expect(Token.IDENT);
        }
        return identifier;
    }

	String parseQualified() throws Exception {
		String qualified = this.parseIdentifier().name;
		while(this.token == Token.Dot) {
			this.next();
			qualified += "." + this.parseIdentifier().name;
		}
		return qualified;
	}
}
