package com.github.panda_io.micro_panda.scanner;

import java.util.*;

public class OperatorNode {
    Map<Byte, OperatorNode> children;
    Token token;

    static OperatorNode root;

    static class Operator {
        Token token;
        int length;

        Operator(Token token, int length) {
            this.token = token;
            this.length = length;
        }
    }

    OperatorNode() {
        this.children = new HashMap<>();
        this.token = Token.ILLEGAL;
    }

    static Operator readOperator(byte[] bytes) {
        return root.find(bytes);
    }

    void insert(String operator) {
        this.insertOperator(operator.getBytes(), 0);
    }
    
    void insertOperator(byte[] operator, int position) {
        if (position < operator.length) {
            Byte character = operator[position];
            if (!this.children.containsKey(character)) {
                OperatorNode node = new OperatorNode();
                this.children.put(character, node);
            }
            position++;
            this.children.get(character).insertOperator(operator, position);
        } else {
            this.token = Token.readToken(new String(operator));
        }
    }   

    Operator find(byte[] operator) {
        return this.findOperator(operator, 0);
    }
    
    Operator findOperator(byte[] operator, int offset) {
        if (this.children.containsKey(operator[offset])) {
            offset++;
            if (offset < operator.length) {
                return this.children.get(operator[offset]).findOperator(operator, offset);
            }
        } else if (offset > 0) {
            //return ReadToken(string(bytes[:offset])), offset
            //TO-DO
            return null;
        }
        return new Operator(Token.ILLEGAL, 1);
    }
}

