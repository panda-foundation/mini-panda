package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/parser"
	"github.com/panda-io/micro-panda/target/llvm"
	"github.com/panda-io/micro-panda/token"
)

type Compiler struct {
	parser  *parser.Parser
	fileset *token.FileSet
	program *ast.Program

	parsedFolder map[string]bool
}

func NewCompiler(flags []string) *Compiler {
	p := ast.NewProgram()
	return &Compiler{
		parser:  parser.NewParser(flags, p),
		fileset: &token.FileSet{},
		program: p,

		parsedFolder: make(map[string]bool),
	}
}

func (c *Compiler) Compile() {
	c.ParseFolder("./")
	if c.Validate() {
		c.GenerateIR("./main")
	}
}

func (c *Compiler) ParseFile(file string) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	f := c.fileset.AddFile(file, len(b))
	c.parser.ParseFile(f, b)
}

func (c *Compiler) ParseFolder(folder string) {
	if c.parsedFolder[folder] {
		return
	}
	folderInfo, err := os.Open(folder)
	if err != nil {
		panic(err)
	}
	list, err := folderInfo.Readdir(-1)
	if err != nil {
		panic(err)
	}
	for _, f := range list {
		if f.IsDir() {
			c.ParseFolder(filepath.Join(folder, f.Name()))
		} else {
			if strings.HasSuffix(f.Name(), ".mpd") {
				filename := filepath.Join(folder, f.Name())
				c.ParseFile(filename)
			}
		}
	}
}

func (c *Compiler) Validate() bool {
	c.program.Validate()
	return !c.program.PrintErrors()
}

func (c *Compiler) GenerateIR(file string) {
	p := llvm.NewProgram()
	content := p.GenerateIR(c.program)
	if err := ioutil.WriteFile(file+".ll", []byte(content), 0644); err != nil {
		panic(err)
	}
	cmd := exec.Command("opt-10", "-o", file+".opt.ll", "-S", "--O2", file+".ll")
	if out, err := cmd.CombinedOutput(); err != nil {
		if err != nil {
			fmt.Println(string(out))
		}
		return
	}
	cmd = exec.Command("llc-10", "-filetype=obj", "-o", file+".o", file+".opt.ll")
	if out, err := cmd.CombinedOutput(); err != nil {
		if err != nil {
			fmt.Println(string(out))
		}
		return
	}
	cmd = exec.Command("clang", "-o", file, file+".o")
	if out, err := cmd.CombinedOutput(); err != nil {
		if err != nil {
			fmt.Println(string(out))
		}
		return
	}
}
