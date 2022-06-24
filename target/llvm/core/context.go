package core

import (
	"github.com/panda-io/micro-panda/ir/core"
)

type Context interface {
	NewContext() Context
	AddObject(name string, value core.Value)
	FindObject(name string) core.Value
	AutoLoad(value core.Value) core.Value
}
