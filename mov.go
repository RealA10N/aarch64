package aarch64

import "fmt"

type MovShift uint8

func (m MovShift) Validate() error {
	if m > 3 {
		return fmt.Errorf("invalid shift value %d", m)
	}
	return nil
}

func (m MovShift) Binary() uint32 {
	return uint32(m)
}

type movz uint32

func MOVZ(dest GPRegister, imm Immediate16, shift MovShift) (movz, error) {
	inst := (0b110100101 << 23) | (shift.Binary() << 21) | (imm.Binary() << 5) | (dest.Binary())
	return movz(inst), nil
}

type movk uint32

func MOVK(dest GPRegister, imm Immediate16, shift MovShift) (movk, error) {
	inst := (0b111100101 << 23) | (shift.Binary() << 21) | (imm.Binary() << 5) | (dest.Binary())
	return movk(inst), nil
}

type movn uint32

func MOVN(dest GPRegister, imm Immediate16, shift MovShift) (movn, error) {
	inst := (0b100100101 << 23) | (shift.Binary() << 21) | (imm.Binary() << 5) | (dest.Binary())
	return movn(inst), nil
}
