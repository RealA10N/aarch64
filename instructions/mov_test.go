package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
	"github.com/stretchr/testify/assert"
)

func TestMovz1(t *testing.T) {
	expected := assembleInstruction(t, "MOVZ x0, #0x1234, LSL #32")
	actual := instructions.MOVZ(
		registers.X0,
		immediates.Immediate16(0x1234),
		instructions.MovShift32,
	)
	assert.Equal(t, expected, uint32(actual))
}

func TestMovz2(t *testing.T) {
	expected := assembleInstruction(t, "MOVZ x30, #0xFFFF")
	actual := instructions.MOVZ(
		registers.X30,
		immediates.Immediate16(0xFFFF),
		instructions.MovShift0,
	)
	assert.Equal(t, expected, uint32(actual))
}

func TestMovk1(t *testing.T) {
	expected := assembleInstruction(t, "MOVK X10, #0, LSL #0x30")
	actual := instructions.MOVK(
		registers.X10,
		immediates.Immediate16(0),
		instructions.MovShift48,
	)
	assert.Equal(t, expected, uint32(actual))
}

func TestMovk2(t *testing.T) {
	expected := assembleInstruction(t, "MOVK X29, #0xFFFE, LSL #16")
	actual := instructions.MOVK(
		registers.X29,
		immediates.Immediate16(0xFFFE),
		instructions.MovShift16,
	)
	assert.Equal(t, expected, uint32(actual))
}

func TestMovn1(t *testing.T) {
	expected := assembleInstruction(t, "MOVN X2, #0xBEEF, LSL #32")
	actual := instructions.MOVN(
		registers.X2,
		immediates.Immediate16(0xBEEF),
		instructions.MovShift32,
	)
	assert.Equal(t, expected, uint32(actual))
}

func TestMovn2(t *testing.T) {
	expected := assembleInstruction(t, "MOVN X21, #0x0, LSL #0")
	actual := instructions.MOVN(
		registers.X21,
		immediates.Immediate16(0),
		instructions.MovShift0,
	)
	assert.Equal(t, expected, uint32(actual))
}
