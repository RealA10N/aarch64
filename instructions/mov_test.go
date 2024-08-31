package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestMovz1(t *testing.T) {
	AssertExpectedInstruction(t, "MOVZ X0, #1234, LSL #32", instructions.MOVZ(
		registers.X0,
		immediates.Immediate16(1234),
		instructions.MovShift32,
	))
}

func TestMovz2(t *testing.T) {
	AssertExpectedInstruction(t, "MOVZ X30, #65535", instructions.MOVZ(
		registers.X30,
		immediates.Immediate16(0xFFFF),
		instructions.MovShift0,
	))
}

func TestMovk1(t *testing.T) {
	AssertExpectedInstruction(t, "MOVK X10, #0, LSL #0x30", instructions.MOVK(
		registers.X10,
		immediates.Immediate16(0),
		instructions.MovShift48,
	))
}

func TestMovk2(t *testing.T) {
	AssertExpectedInstruction(t, "MOVK X29, #0x65534, LSL #16", instructions.MOVK(
		registers.X29,
		immediates.Immediate16(0xFFFE),
		instructions.MovShift16,
	))
}

func TestMovn1(t *testing.T) {
	AssertExpectedInstruction(t, "MOVN X2, #48879, LSL #32", instructions.MOVN(
		registers.X2,
		immediates.Immediate16(0xBEEF),
		instructions.MovShift32,
	))
}

func TestMovn2(t *testing.T) {
	AssertExpectedInstruction(t, "MOVN X21, #0, LSL #0", instructions.MOVN(
		registers.X21,
		immediates.Immediate16(0),
		instructions.MovShift0,
	))
}
