package com.github.panda_io.micro_panda;

import com.github.panda_io.micro_panda.parser.*;
import com.github.panda_io.micro_panda.ast.expression.*;

public class App {
    public static void main(String[] args) {
        Parser parser = new Parser(null);
        Expression expression = parser.parseExpression("1 + 1");
        System.out.println(expression);
    }
}
