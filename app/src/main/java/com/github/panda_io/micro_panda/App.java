package com.github.panda_io.micro_panda;

import com.github.panda_io.micro_panda.scanner.File;
public class App {
    public String getGreeting() {
        return "Hello World!";
    }

    public static void main(String[] args) {
        System.out.println(new App().getGreeting());
        File file = new File("file1", 999);
    }
}
