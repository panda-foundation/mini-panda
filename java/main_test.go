package main

import (
	"testing"
)

func TestMiniPanda(t *testing.T) {
	c := NewCompiler(nil)
	c.ParseFile("./micro-panda/libc/io.mpd")
	c.ParseFile("./micro-panda/libc/memory.mpd")
	c.ParseFile("./micro-panda/llvm/memory.mpd")
	c.ParseFile("./micro-panda/test/test.mpd")
	c.ParseFile("./micro-panda/test/expression.mpd")
	c.ParseFile("./micro-panda/test/function.mpd")
	c.ParseFile("./micro-panda/test/statement.mpd")
	c.ParseFile("./micro-panda/test/struct.mpd")
	c.ParseFile("./micro-panda/main.mpd")
	if c.Validate() {
		c.GenerateIR("./micro-panda/main")
	}
	t.Fail()
}
