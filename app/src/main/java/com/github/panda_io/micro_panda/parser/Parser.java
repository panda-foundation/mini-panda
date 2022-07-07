package com.github.panda_io.micro_panda.parser;

import java.io.FileInputStream;
import java.util.Set;

import com.github.panda_io.micro_panda.scanner.*;
import com.github.panda_io.micro_panda.ast.Program;
import com.github.panda_io.micro_panda.ast.Module;
import com.github.panda_io.micro_panda.ast.expression.Expression;
import com.github.panda_io.micro_panda.ast.statement.Statement;

public class Parser {
    Program program;
    Scanner scanner;
    FileSet files;
    Context context;

    public Parser(Set<String> flags) {
        this.program = new Program();
        this.scanner = new Scanner(flags);
        this.files = new FileSet();
        this.context = new Context(this.program, this.scanner);
    }

    public void parseFile(String path) {
        try {
            java.io.File input = new java.io.File(path);
            FileInputStream fs = new FileInputStream(input);
            int size = (int) input.length();
            byte[] source = new byte[size];
            fs.read(source);
            fs.close();

            File file = new File(path, size);
            this.loadSource(file, source);
            Module module = ModuleParser.parseModule(context, file);
            this.program.addModule(path, module);
        } catch (Exception e) {
            System.out.printf("parse file \"%s\" failed:\n", path);
            System.out.println(e.getMessage());
        } finally {
            this.unloadSource();
        }
    }

    public Module parseBytes(byte[] source) {
        try {
            File file = new File("<input>", source.length);
            this.loadSource(file, source);
            Module module = ModuleParser.parseModule(context, file);
            return module;
        } catch (Exception e) {
            System.out.println("parse bytes failed:");
            System.out.println(e.getMessage());
        } finally {
            this.unloadSource();
        }
        return null;
    }

    public Expression parseExpression(String source) {
        try {
            byte[] bytes = source.getBytes();
            File file = new File("<input>", bytes.length);
            this.loadSource(file, bytes);
            Expression expression = ExpressionParser.parseExpression(this.context);
            return expression;
        } catch (Exception e) {
            System.out.println("parse expression failed:");
            System.out.println(e.getMessage());
        } finally {
            this.unloadSource();
        }
        return null;
    }

    public Statement parseStatements(String source) {
        try {
            byte[] bytes = source.getBytes();
            File file = new File("<input>", bytes.length);
            this.loadSource(file, bytes);
            Statement statement = StatementParser.parseStatement(this.context);
            return statement;
        } catch (Exception e) {
            System.out.println("parse statement failed:");
            System.out.println(e.getMessage());
        } finally {
            this.unloadSource();
        }
        return null;
    }

    void loadSource(File file, byte[] source) throws Exception {
        this.scanner.loadSource(file, source);
        this.scanner.scan();
    }

    void unloadSource() {
        try {
            this.scanner.close();
        } catch (Exception e) {
        }
    }
}
