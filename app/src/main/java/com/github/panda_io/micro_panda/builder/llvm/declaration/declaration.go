package llvm

type Declaration interface {
	Declaration()
}

type DeclarationBase struct {
	Qualified string
}

func (*DeclarationBase) Declaration() {
}
