package com.github.panda_io.micro_panda;

import com.github.panda_io.micro_panda.scanner.*;

public class App {
    public static void main(String[] args) {
        String message = "Hello Micro Panda";
        byte[] data = message.getBytes();
        File file = new File("<in memory>", data.length);
        Scanner scanner = new Scanner(null);
        try {
            scanner.loadSource(file, data);
            while (scanner.token != Token.EOF) {
                scanner.scan();
                if (scanner.token != Token.EOF) {
                    System.out.printf("scan: %s \n", scanner.literal);
                }
            }

        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
    }
}
