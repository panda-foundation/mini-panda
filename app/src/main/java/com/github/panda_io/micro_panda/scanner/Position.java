package com.github.panda_io.micro_panda.scanner;

public class Position {
    File file;
    int offset;

    public Position(File file, int offset) {
        this.file = file;
        this.offset = offset;
    }

    public String getLocation() {
        File.Location location = this.file.getLocation(this.offset);
        return String.format("%s:%d:%d", this.file.name, location.line, location.column);
    }

    public int getGlobalOffset() {
        return this.file.base + this.offset;
    }
}
