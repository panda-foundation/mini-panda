package com.github.panda_io.micro_panda.parser;

import com.github.panda_io.micro_panda.ast.Program;
import com.github.panda_io.micro_panda.ast.expression.Identifier;
import com.github.panda_io.micro_panda.scanner.Scanner;
import com.github.panda_io.micro_panda.scanner.Token;

public class Context {
    Program program;
    Scanner scanner;

    public Context(Program program, Scanner scanner) {
        this.program = program;
        this.scanner = scanner;
    }

    void expect(Token token) throws Exception {
        if (this.scanner.token != token) {
            this.expectedError(this.scanner.position, String.format("'%s'", token.toString()));
        }
        this.scanner.scan();
    }

    void expectedError(int position, String expect) {
        expect = "expected " + expect;
        if (position == this.scanner.position) {
            if (this.scanner.token == Token.Semi && this.scanner.literal == "\n") {
                expect += ", but found newline";
            } else if (this.scanner.token.isLiteral()){
                expect += ", but found " + this.scanner.literal;
            } else {
                expect += ", but found '" + this.scanner.token.toString() + "'";
            }
        }
        this.program.addError(position, expect);
    }

	void addError(int offset, String message) {
        this.program.addError(offset, message);
    }

    Identifier  parseIdentifier() throws Exception {
        Identifier identifier = new Identifier();
        identifier.setOffset(this.scanner.position);
        if (this.scanner.token == Token.IDENT) {
            identifier.name = this.scanner.literal;
            this.scanner.scan();
        } else {
            this.expect(Token.IDENT);
        }
        return identifier;
    }

	String parseQualified() throws Exception {
		String qualified = this.parseIdentifier().name;
		while(this.scanner.token == Token.Dot) {
			this.scanner.scan();
			qualified += "." + this.parseIdentifier().name;
		}
		return qualified;
	}
}
