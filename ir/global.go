package ir

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/ir/constant"
	"github.com/panda-io/micro-panda/ir/core"
	"github.com/panda-io/micro-panda/ir/types"
)

type Global struct {
	core.GlobalIdent
	Immutable   bool
	ContentType core.Type
	Init        constant.Constant
	Typ         *types.PointerType
}

func NewGlobal(name string, contentType core.Type) *Global {
	global := &Global{ContentType: contentType}
	global.SetName(name)
	global.Type()
	return global
}

func NewGlobalDef(name string, init constant.Constant) *Global {
	global := &Global{ContentType: init.Type(), Init: init}
	global.SetName(name)
	global.Type()
	return global
}

func (g *Global) String() string {
	return fmt.Sprintf("%s %s", g.Type(), g.Ident())
}

func (g *Global) Type() core.Type {
	if g.Typ == nil {
		g.Typ = types.NewPointerType(g.ContentType)
	}
	return g.Typ
}

func (g *Global) WriteIR(w io.Writer) error {
	declaration := "global"
	if g.Immutable {
		declaration = "constant"
	}
	_, err := fmt.Fprintf(w, "%s = %s %s", g.Ident(), declaration, g.ContentType)
	if err != nil {
		return err
	}
	if g.Init != nil {
		_, err = fmt.Fprintf(w, " %s", g.Init.Ident())
	}
	return err
}

func (g *Global) IsConstant() {}
