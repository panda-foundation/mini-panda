package com.github.panda_io.micro_panda.ast;

public abstract class Node {
    int offset;

    public int getOffset() {
        return this.offset;
    }

    public void setOffset(int offset) {
        this.offset = offset;
    }
}
