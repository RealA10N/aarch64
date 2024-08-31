// An implementation of the MOVZ, MOVK and MOVN
// Those are all of the explicit MOV instruction variants of an immediate to a
// general purpose register in AArch64. The "MOV" instruction itself is an alias
// to MOVZ with a shift of 0, and (when possible) a MOVN instruction with a high
// immediate value that can be encoded using 16 bits and shifts.
package instructions

import (
	"fmt"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/registers"
)

type MovShift uint8

const (
	MovShift0 MovShift = iota
	MovShift16
	MovShift32
	MovShift48
)

func (m MovShift) Validate() error {
	if m > 3 {
		return fmt.Errorf("invalid shift value %d", m)
	}
	return nil
}

func (m MovShift) Binary() uint32 {
	return uint32(m)
}

func (m MovShift) String() string {
	return fmt.Sprintf("LSL #%d", m*16)
}

func movBuildRaw(magic uint32, dest registers.GPRegister, imm immediates.Immediate16, shift MovShift) uint32 {
	return (magic << 23) | (shift.Binary() << 21) | (imm.Binary() << 5) | (dest.Binary())
}

func movString(name string, dest registers.GPRegister, imm immediates.Immediate16, shift MovShift) string {
	s := fmt.Sprintf("%s %s, %s", name, dest, imm)
	if shift != 0 {
		s += fmt.Sprintf(", %s", shift)
	}
	return s
}

func movExtractDetails(m uint32) (uint32, registers.GPRegister, immediates.Immediate16, MovShift) {
	magic := m >> 23
	shift := MovShift((m >> 21) & 0b11)
	imm := immediates.Immediate16((m >> 5) & 0xFFFF)
	dest := registers.GPRegister(m & 0b11111)
	return magic, dest, imm, shift
}

// MARK: MOVZ

type movz uint32

func MOVZ(dest registers.GPRegister, imm immediates.Immediate16, shift MovShift) movz {
	return movz(movBuildRaw(0b110100101, dest, imm, shift))
}

func (m movz) String() string {
	_, dest, imm, shift := movExtractDetails(uint32(m))
	return movString("MOVZ", dest, imm, shift)
}

// MARK: MOVK

type movk uint32

func MOVK(dest registers.GPRegister, imm immediates.Immediate16, shift MovShift) movk {
	return movk(movBuildRaw(0b111100101, dest, imm, shift))
}

func (m movk) String() string {
	_, dest, imm, shift := movExtractDetails(uint32(m))
	return movString("MOVK", dest, imm, shift)
}

// MARK: MOVN

type movn uint32

func MOVN(dest registers.GPRegister, imm immediates.Immediate16, shift MovShift) movn {
	return movn(movBuildRaw(0b100100101, dest, imm, shift))
}

func (m movn) String() string {
	_, dest, imm, shift := movExtractDetails(uint32(m))
	return movString("MOVN", dest, imm, shift)
}
