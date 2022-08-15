package com.github.panda_io.micro_panda.parser;

public class PreprocessorParser {
    /*
	//#if #else #elif #end
	preprocessorIf     = "if"
	preprocessorElse   = "else"
	preprocessorElseIf = "elif"
	preprocessorEnd    = "end"

	// TO-DO add logical operator support for preprocessor
	// TO-DO separate preprocessor
	// () == != ! && ||
	// #macro
	*/

    //Binary:  || && < <= > >= == !=
    //Unary: !
    // ()
    /*

func (s *Scanner) scanPreprossesor() (int, token.Token, string) {
	//#if #else #elif #end
	if !s.isLetter(s.char) {
		s.error(s.offset, "unexpected identifier")
	}
	literal := s.scanIdentifier()
	if literal == preprocessorIf {
		s.preprocessorLevel++
		s.preprocessorStack = append(s.preprocessorStack, &preprocessor{
			currentBlock: preprocessorIf,
			satisfied:    false,
		})

		result := s.scanPreprossesorExpression()
		if result {
			s.preprocessorStack[s.preprocessorLevel-1].satisfied = true
		} else {
			s.skipPreprossesor()
		}
	} else if literal == preprocessorElseIf {
		if s.preprocessorLevel == 0 || s.preprocessorStack[s.preprocessorLevel-1].currentBlock == preprocessorElse {
			s.error(s.offset, "unexpected #elif")
		} else if s.preprocessorStack[s.preprocessorLevel-1].satisfied {
			s.skipPreprossesor()
		} else {
			if s.scanPreprossesorExpression() {
				s.preprocessorStack[s.preprocessorLevel-1].satisfied = true
			} else {
				s.skipPreprossesor()
			}
		}
		s.preprocessorStack[s.preprocessorLevel-1].currentBlock = preprocessorElseIf
	} else if literal == preprocessorElse {
		if s.preprocessorLevel == 0 || s.preprocessorStack[s.preprocessorLevel-1].currentBlock == preprocessorElse {
			s.error(s.offset, "unexpected #else")
		} else if s.preprocessorStack[s.preprocessorLevel-1].satisfied {
			s.skipPreprossesor()
		}
		s.preprocessorStack[s.preprocessorLevel-1].currentBlock = preprocessorElse
	} else if literal == preprocessorEnd {
		if s.preprocessorLevel == 0 {
			s.error(s.offset, "unexpected #end")
		}
		s.preprocessorLevel--
		s.preprocessorStack = s.preprocessorStack[:s.preprocessorLevel]
	} else {
		s.error(s.offset, "unexpected preprocessor: "+literal)
	}

	return s.Scan()
}

func (s *Scanner) scanPreprossesorExpression() bool {
	for s.char == ' ' || s.char == '\t' {
		s.next()
	}

	if s.isLetter(s.char) {
		flag := s.scanIdentifier()

		for s.char == ' ' || s.char == '\t' || s.char == '\r' {
			s.next()
		}
		if s.char != '\n' {
			s.error(s.offset, "unexpected: "+string(s.char))
		}

		result := false
		if _, ok := s.flags[flag]; ok {
			result = true
		}
		return result
	}

	s.error(s.offset, "unexpected: "+string(s.char))
	return false
}

func (s *Scanner) skipPreprossesor() {
	level := s.preprocessorLevel
	for {
		for s.char != eof && s.char != '#' {
			s.next()
		}
		if s.char == eof {
			s.error(s.offset, "preprocessor not terminated, expecting #end")
		}
		offset := s.offset
		readOffset := s.readOffset
		s.next()
		if s.isLetter(s.char) {
			literal := s.scanIdentifier()

			if literal == preprocessorIf {
				s.preprocessorLevel++
				s.preprocessorStack = append(s.preprocessorStack, &preprocessor{
					currentBlock: preprocessorIf,
					satisfied:    false,
				})
			} else if literal == preprocessorElseIf {
				if s.preprocessorLevel == level {
					s.offset = offset
					s.readOffset = readOffset
					s.char = '#'
					break
				}
				if s.preprocessorStack[s.preprocessorLevel-1].currentBlock == preprocessorElse {
					s.error(s.offset, "unexpected #elif")
				}
				s.preprocessorStack[s.preprocessorLevel-1].currentBlock = preprocessorElseIf
			} else if literal == preprocessorElse {
				if s.preprocessorLevel == level {
					s.offset = offset
					s.readOffset = readOffset
					s.char = '#'
					break
				}
				if s.preprocessorStack[s.preprocessorLevel-1].currentBlock == preprocessorElse {
					s.error(s.offset, "unexpected #else")
				}
				s.preprocessorStack[s.preprocessorLevel-1].currentBlock = preprocessorElse
			} else if literal == preprocessorEnd {
				if s.preprocessorLevel == level {
					s.offset = offset
					s.readOffset = readOffset
					s.char = '#'
					break
				}
				s.preprocessorLevel--
				s.preprocessorStack = s.preprocessorStack[:s.preprocessorLevel]
			} else {
				s.error(s.offset, "unexpected preprocessor: "+literal)
			}
		} else {
			s.error(s.offset, "expected identifier")
		}
	}
}
*/
}
