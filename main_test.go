package main

import (
	"testing"
)

func TestMiniPanda(t *testing.T) {
	c := NewCompiler(nil)
	c.ParseFile("./micro-panda/console/write.mpd")
	//c.ParseFile("./micro-panda/test/test.mpd")
	//c.ParseFile("./micro-panda/test/expression.mpd")
	//c.ParseFile("./micro-panda/test/function.mpd")
	//c.ParseFile("./micro-panda/test/statement.mpd")
	//c.ParseFile("./micro-panda/test/struct.mpd")
	//c.ParseFile("./micro-panda/main.mpd")
	c.ParseFile("./micro-panda/tmp.mpd")
	if c.Validate() {
		//c.GenerateIR("./micro-panda/main")
	}
	t.Fail()
}
