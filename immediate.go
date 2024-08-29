package aarch64

import "fmt"

type Immediate16 uint16

func (Immediate16) Validate() error {
	return nil
}

func (imm Immediate16) Binary() uint32 {
	return uint32(imm)
}

func (imm Immediate16) String() string {
	return fmt.Sprintf("#%d", imm)
}
