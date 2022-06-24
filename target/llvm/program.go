package llvm

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"sort"

	"github.com/panda-io/micro-panda/ast"
	"github.com/panda-io/micro-panda/target/llvm/declaration"
	"github.com/panda-io/micro-panda/target/llvm/ir"
	"github.com/panda-io/micro-panda/target/llvm/ir/constant"
)

type Program struct {
	Program      *ir.Program
	Strings      map[string]constant.Constant
	Declarations map[string]declaration.Declaration
}

func NewProgram() *Program {
	p := &Program{}
	p.Reset()
	return p
}

func (p *Program) AddString(value string) constant.Constant {
	bytes := []byte(value)
	bytes = append(bytes, 0)
	hash := fmt.Sprintf("%x", md5.Sum(bytes))
	if v, ok := p.Strings[hash]; ok {
		return v
	}
	array := constant.NewCharArray(bytes)
	s := p.Program.NewGlobalDef("string."+hash, array)
	s.Immutable = true
	p.Strings[hash] = s
	return s
}

func (p *Program) AddDeclaration(qualified string, d declaration.Declaration) {
	p.Declarations[qualified] = d
}

func (p *Program) FindDeclaration(qualified string) declaration.Declaration {
	return p.Declarations[qualified]
}

func (p *Program) Reset() {
	p.Program = ir.NewProgram()
	p.Strings = make(map[string]constant.Constant)
	p.Declarations = make(map[string]declaration.Declaration)
}

func (p *Program) GenerateIR(program *ast.Program) []byte {
	var keys Keys
	for key := range program.Modules {
		keys = append(keys, key)
	}
	sort.Sort(keys)

	var modules []*ast.Module
	for _, key := range keys {
		modules = append(modules, program.Modules[key])
	}

	// first pass (generate declarations)
	for _, m := range modules {
		program.Module = m
		for _, f := range m.Functions {
			ff := &declaration.Function{}
			ff.GenerateDefineIR(p, f)
			p.AddDeclaration(f.QualifiedName(), ff)
		}
		for _, e := range m.Enums {
			ee := &declaration.Enum{}
			ee.GenerateIR(p, e)
			p.AddDeclaration(e.QualifiedName(), ee)
		}
		for _, s := range m.Structs {
			ss := &declaration.Struct{}
			ss.GenerateDefineIR(p, s)
			p.AddDeclaration(s.QualifiedName(), ss)
		}
	}

	// second pass (generate functions)
	for _, m := range modules {
		program.Module = m
		for _, v := range m.Variables {
			vv := &declaration.Variable{}
			vv.GenerateIR(p, v)
			p.AddDeclaration(v.QualifiedName(), vv)
		}
		for _, f := range m.Functions {
			ff := p.FindDeclaration(f.QualifiedName()).(*declaration.Function)
			ff.GenerateIR(p, f)
		}
		for _, s := range m.Structs {
			ss := p.FindDeclaration(s.QualifiedName()).(*declaration.Struct)
			ss.GenerateIR(p, s)
		}
	}

	buf := &bytes.Buffer{}
	err := p.Program.WriteIR(buf)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

type Keys []string

func (list Keys) Len() int      { return len(list) }
func (list Keys) Swap(i, j int) { list[i], list[j] = list[j], list[i] }
func (list Keys) Less(i, j int) bool {
	return list[i] < list[j]
}
