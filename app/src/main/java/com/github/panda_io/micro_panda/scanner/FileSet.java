package com.github.panda_io.micro_panda.scanner;

import java.util.*;

public class FileSet {
    List<File> files;
    int base;

    public FileSet() {
        this.files = new ArrayList<>();
        this.base = 0;
    }

    public File addFile(String fileName, int size) {
        for (File file : this.files) {
            if (file.name.equals(fileName)) {
                throw new RuntimeException(String.format("file %s already added \n", fileName));
            }
        }
        File file = new File(fileName, size);
        file.base = this.base;
        this.base += size + 1;
        this.files.add(file);
        return file;
    }

    public void updateFile(String fileName, int size) {
        boolean found = false;
        for (File file : this.files) {
            if (file.name.equals(fileName)) {
                found = true;
                file.size = size;
            }
        }
        if (found) {
            this.base = 0;
            for (File file : this.files) {
                file.base = this.base;
                this.base += file.size + 1;
            }
        }
    }

    public File getFile(int globalOffset) {
        for (File file : this.files) {
            if (globalOffset <= file.base + file.size) {
                return file;
            }
        }
        return null;
    }

    public Position getPosition(int globalOffset) {
        File file = this.getFile(globalOffset);
        if (file != null) {
            return file.getPosition(globalOffset - file.base);
        }
        return null;
    }
}
