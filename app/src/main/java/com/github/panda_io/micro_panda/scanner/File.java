package com.github.panda_io.micro_panda.scanner;

import java.util.*;

public class File {
    String name;
    int size;
    int base;
    List<Integer> lines;

    static class Location {
        int line;
        int column;

        Location(int line, int column) {
            this.line = line;
            this.column = column;
        }
    }

    public File(String name, int size) {
        this.name = name;
        this.size = size;
        this.lines = new ArrayList<>();
        this.lines.add(0);
    }

    public String filename() {
        return this.name;
    }

    public void addLine(int offset) {
        this.lines.add(offset);
    }

    public int linesSize() {
        return this.lines.size();
    }

    public Position getPosition(int offset) {
        return new Position(this, offset);
    }

    Location getLocation(int offset) {
        int i = 0;
        int j = this.lines.size();
        while (i < j) {
            int k = i + (j - i) / 2;
            if (this.lines.get(k) <= offset) {
                i = k + 1;
            } else {
                j = k;
            }
        }
        Location location = new Location(0, 0);
        i = i - 1;
        if (i >= 0) {
            location.line = i + 1;
            location.column = offset - this.lines.get(i) + 1;
        }
        return location;
    }
}
