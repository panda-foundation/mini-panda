package com.github.panda_io.micro_panda.scanner;

import java.io.*;
import java.nio.charset.StandardCharsets;
import java.util.Arrays;
import java.util.HashSet;
import java.util.Set;

public class Scanner {
	public static final int BOM = 0xFEFF;
	public static final int EOF = -1;

	public int position;
	public Token token;
	public String literal;

	Set<String> flags;

	int character;
	int offset;
	int readOffset;

	File file;
	PushbackReader input;
	byte[] source;

	public Scanner(Set<String> flags) {
		this.flags = flags;
		if (this.flags == null) {
			this.flags = new HashSet<>();
		}
	}

	public void loadSource(File file, byte[] source) throws Exception {
		this.file = file;
		this.source = source;
		this.close();
		this.input = new PushbackReader(
				new InputStreamReader(new ByteArrayInputStream(source), StandardCharsets.UTF_8));
		this.character = 0;
		this.offset = 0;
		this.readOffset = 0;
		this.position = 0;
		this.token = Token.ILLEGAL;
		this.literal = "";
		this.next();
		if (this.character == BOM) {
			this.next();
		}
	}

	public void close() {
		if (this.input != null) {
			try {
				this.input.close();
			} catch (Exception e) {
				System.out.printf("close input stream failed: %s \n", e.getMessage());
			}
		}
	}

	public void scan() throws Exception {
		while (this.character == ' ' || this.character == '\t' || this.character == '\n' || this.character == '\r') {
			this.next();
		}
		this.position = this.offset;
		this.token = Token.ILLEGAL;

		if (this.isLetter(this.character)) {
			this.literal = this.scanIdentifier();
			this.token = Token.readToken(this.literal);
		} else if (this.isDecimal(this.character) || (this.character == '.' && this.isDecimal(this.peek()))) {
			this.literal = this.scanNumber();
		} else {
			int character = this.character;
			this.next();
			switch (character) {
				case EOF:
					this.token = Token.EOF;
					/*
					 * if s.preprocessorLevel > 0 {
					 * s.error(s.offset, "preprocessor not terminated, expecting #end")
					 * }
					 */
					break;

				case '\'':
					this.token = Token.CHAR;
					this.literal = this.scanChar();
					break;

				case '"':
					this.token = Token.STRING;
					this.literal = this.scanString();
					break;

				case '`':
					this.token = Token.STRING;
					this.literal = this.scanRawString();
					break;

				case '/':
					if (this.character == '/' || this.character == '*') {
						this.scanComment();
						this.scan();
						break;
					}
					this.literal = this.scanOperators();
					break;

				case '@':
					this.token = Token.META;
					this.literal = "@";
					break;

				case ';':
					this.token = Token.Semi;
					this.literal = ";";
					break;

				// case '#':
				// return s.scanPreprossesor()

				default:
					this.literal = this.scanOperators();
					if (this.token == Token.ILLEGAL) {
						throw new RuntimeException(
								String.format("invalid token:\n%s", this.file.getPosition(this.offset).string()));
					}
			}
		}
	}

	void next() throws Exception {
		this.character = this.input.read();
		if (this.character != EOF) {
			this.offset = this.readOffset;
			byte first = (byte) (this.character & 255);
			if (first < 0x80) {
				this.readOffset += 1;
			} else if (first < 0xe0) {
				this.readOffset += 2;
			} else if (first < 0xf0) {
				this.readOffset += 3;
			} else {
				this.readOffset += 4;
			}
			if (this.character == '\n') {
				this.file.addLine(this.offset);
			}
		} else {
			this.offset = this.file.size;
		}
	}

	int peek() throws Exception {
		int character = this.input.read();
		this.input.unread(character);
		return character;
	}

	String scanIdentifier() throws Exception {
		int start = this.offset;
		while (this.isLetter(this.character) || this.isDecimal(this.character)) {
			this.next();
		}
		return new String(this.source, start, this.offset - start);
	}

	String scanComment() throws Exception {
		int start = this.offset - 1;
		if (this.character == '/') {
			// -style comment
			this.next();
			while (this.character != '\n' && this.character >= 0) {
				this.next();

			}
		} else {
			/*-style comment */
			boolean terminated = false;
			this.next();
			while (this.character >= 0) {
				int character = this.character;
				this.next();
				if (character == '*' && this.character == '/') {
					this.next();
					terminated = true;
					break;
				}
			}
			if (!terminated) {
				throw new RuntimeException(
						String.format("comment not terminated:\n%s", this.file.getPosition(this.offset).string()));
			}
		}
		return new String(this.source, start, this.offset - start, StandardCharsets.UTF_8);
	}

	String scanNumber() throws Exception {
		int start = this.offset;
		this.token = Token.INT;
		if (this.character != '.') {
			if (this.character == '0') {
				this.next();
				if (this.character != '.') {
					int base = 10;
					switch (this.lower(this.character)) {
						case 'x':
							base = 16;
							break;
						case 'b':
							base = 2;
							break;
						case 'o':
							base = 8;
							break;
						default:
							if (this.isDecimal(this.character)) {
								this.token = Token.ILLEGAL;
								throw new RuntimeException(String.format("invalid integer:\n%s",
										this.file.getPosition(this.offset).string()));
							} else {
								return "0";
							}
					}
					if (this.token != Token.ILLEGAL) {
						this.next();
						this.scanDigits(base);
						if (this.offset - start <= 2) {
							this.token = Token.ILLEGAL;
							throw new RuntimeException(
									String.format("invalid number:\n%s", this.file.getPosition(this.offset).string()));
						}
						if (this.character == '.') {
							this.token = Token.ILLEGAL;
							throw new RuntimeException(String.format("invalid radix point:\n%s",
									this.file.getPosition(this.offset).string()));
						}
					}
				}
			} else {
				this.scanDigits(10);
			}
		}
		if (this.character == '.') {
			int offset = this.offset;
			this.token = Token.FLOAT;
			this.next();
			this.scanDigits(10);
			if (offset == this.offset - 1) {
				this.token = Token.ILLEGAL;
				throw new RuntimeException(String.format("float has no digits after '.':\n%s",
						this.file.getPosition(this.offset).string()));
			}
		}
		return new String(this.source, start, this.offset - start);
	}

	void scanDigits(int base) throws Exception {
		while (this.digitValue(this.character) < base) {
			this.next();
		}
	}

	String scanString() throws Exception {
		int start = this.offset - 1;
		while (true) {
			int character = this.character;
			if (character == '\n' || character < 0) {
				throw new RuntimeException(String.format("string literal not terminated:\n%s",
						this.file.getPosition(this.offset).string()));
			}
			this.next();
			if (character == '"') {
				break;
			}
			if (character == '\\') {
				this.scanEscape('"');
			}
		}
		return new String(this.source, start, this.offset - start, StandardCharsets.UTF_8);
	}

	String scanChar() throws Exception {
		int start = this.offset - 1;
		int character = this.character;
		if (character == '\n' || character < 0) {
			throw new RuntimeException(
					String.format("char literal not terminated:\n%s", this.file.getPosition(this.offset).string()));
		}
		this.next();
		if (character == '\\') {
			this.scanEscape('\'');
		}
		if (this.character != '\'') {
			throw new RuntimeException(
					String.format("illegal char literal:\n%s", this.file.getPosition(this.offset).string()));
		}
		this.next();
		return new String(this.source, start, this.offset - start);
	}

	String scanRawString() throws Exception {
		int start = this.offset - 1;
		while (true) {
			int character = this.character;
			if (character < 0) {
				throw new RuntimeException(String.format("raw string literal not terminated:\n%s",
						this.file.getPosition(this.offset).string()));
			}
			this.next();
			if (character == '`') {
				break;
			}
		}
		return new String(this.source, start, this.offset - start, StandardCharsets.UTF_8);
	}

	void scanEscape(int quote) throws Exception {
		int start = this.offset;
		int n, base, max = 0;
		if (this.character == quote) {
			this.next();
			return;
		}
		switch (this.character) {
			case 'a':
			case 'b':
			case 'f':
			case 'n':
			case 'r':
			case 't':
			case 'v':
			case '\\':
				this.next();
				return;

			case '0':
			case '1':
			case '2':
			case '3':
			case '4':
			case '5':
			case '6':
			case '7':
				n = 3;
				base = 8;
				max = 255;
				break;

			case 'x':
				this.next();
				n = 2;
				base = 16;
				max = 255;
				break;

			case 'u':
				this.next();
				n = 4;
				base = 16;
				max = Character.MAX_CODE_POINT;
				break;

			case 'U':
				this.next();
				n = 8;
				base = 16;
				max = Character.MAX_CODE_POINT;
				break;

			default:
				if (this.character < 0) {
					throw new RuntimeException(String.format("escape sequence not terminated:\n%s",
							this.file.getPosition(this.offset).string()));
				}
				throw new RuntimeException(
						String.format("unknown escape sequence:\n%s", this.file.getPosition(this.offset).string()));
		}

		int x = 0;
		while (n > 0) {
			int d = this.digitValue(this.character);
			if (d > base) {
				throw new RuntimeException(String.format("illegal character %s in escape sequence:\n%s",
						Character.toString(this.character), this.file.getPosition(this.offset).string()));
			}
			x = x * base + d;
			this.next();
			n--;
		}
		if (x > max || 0xD800 <= x && x < 0xE000) {
			throw new RuntimeException(String.format("escape sequence is invalid Unicode code point:\n%s",
					this.file.getPosition(start).string()));
		}
	}

	String scanOperators() throws Exception {
		int start = this.offset - 1;
		byte[] data = Arrays.copyOfRange(this.source, start, start + 3); // 3 is max length of operator
		OperatorNode.Operator operator = OperatorNode.readOperator(data);
		if (operator.length > 0) {
			for (int i = 1; i < operator.length; i++) {
				this.next();
			}
			this.token = operator.token;
			return new String(this.source, start, this.offset - start);
		}
		return "";
	}

	int digitValue(int character) {
		if ('0' <= character && character <= '9') {
			return character - '0';
		}
		if ('a' <= this.lower(character) && this.lower(character) <= 'f') {
			return this.lower(character) - 'a' + 10;
		}
		return 16; // larger than any legal digit val
	}

	int lower(int character) {
		return ('a' - 'A') | character;
	}

	boolean isLetter(int character) {
		return character == '_' || 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z';
	}

	boolean isDecimal(int character) {
		return '0' <= character && character <= '9';
	}
}
