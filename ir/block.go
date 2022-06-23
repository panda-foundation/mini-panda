package ir

import (
	"fmt"
	"io"

	"github.com/panda-io/micro-panda/ir/core"
	"github.com/panda-io/micro-panda/ir/instruction"
	"github.com/panda-io/micro-panda/ir/types"
)

type Block struct {
	core.LocalIdent
	Insts      []instruction.Instruction
	Terminated bool
}

func NewBlock(name string) *Block {
	block := &Block{}
	block.SetName(name)
	return block
}

func (block *Block) AddInstruction(inst instruction.Instruction) {
	if block.Terminated {
		panic("block already terminated")
	}
	block.Insts = append(block.Insts, inst)
	if _, ok := inst.(instruction.Terminator); ok {
		block.Terminated = true
	}
}

func (block *Block) InsertAlloca(inst *instruction.InstAlloca) {
	if block.Terminated {
		block.Insts = append(block.Insts, block.Insts[len(block.Insts)-1])
		block.Insts[len(block.Insts)-2] = inst
	} else {
		block.Insts = append(block.Insts, inst)
	}
}

func (block *Block) String() string {
	return fmt.Sprintf("%s %s", block.Type(), block.Ident())
}

func (block *Block) Type() core.Type {
	return types.Label
}

func (block *Block) WriteIR(w io.Writer) error {
	name := ""
	if block.LocalName == "" {
		name = core.LabelID(block.LocalID)
	} else {
		name = core.LabelName(block.LocalName)
	}
	_, err := fmt.Fprintf(w, "%s\n", name)
	if err != nil {
		return err
	}
	for _, inst := range block.Insts {
		_, err = w.Write([]byte("\t"))
		if err != nil {
			return err
		}
		err = inst.WriteIR(w)
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}
	return nil
}
