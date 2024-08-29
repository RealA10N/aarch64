package aarch64

type Immediate16 uint16

func (Immediate16) Validate() error {
	return nil
}

func (imm Immediate16) Binary() uint32 {
	return uint32(imm)
}
