package main

import (
	"testing"
)

func TestMiniPanda(t *testing.T) {
	c := NewCompiler(nil)
	c.ParseFile("../mini-panda/libc/io.mpd")
	c.ParseFile("../mini-panda/libc/memory.mpd")
	c.ParseFile("../mini-panda/llvm/memory.mpd")
	c.ParseFile("../mini-panda/test/test.mpd")
	c.ParseFile("../mini-panda/test/expression.mpd")
	c.ParseFile("../mini-panda/test/function.mpd")
	c.ParseFile("../mini-panda/test/statement.mpd")
	c.ParseFile("../mini-panda/test/struct.mpd")
	c.ParseFile("../mini-panda/container/allocator.mpd")
	//c.ParseFile("../mini-panda/container/vector.mpd")
	c.ParseFile("../mini-panda/main.mpd")
	if c.Validate() {
		c.GenerateIR("../mini-panda/main")
	}
	t.Fail()
}
