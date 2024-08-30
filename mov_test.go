package aarch64codegen_test

import (
	"testing"

	"alon.kr/x/aarch64codegen"
	"github.com/stretchr/testify/assert"
)

func TestMovz1(t *testing.T) {
	expected := assembleInstruction(t, "MOVZ x0, #0x1234, LSL #32")
	actual := aarch64codegen.MOVZ(
		aarch64codegen.X0,
		aarch64codegen.Immediate16(0x1234),
		aarch64codegen.MovShift32,
	)
	assert.Equal(t, expected, uint32(actual))
}

func TestMovz2(t *testing.T) {
	expected := assembleInstruction(t, "MOVZ x30, #0xFFFF")
	actual := aarch64codegen.MOVZ(
		aarch64codegen.X30,
		aarch64codegen.Immediate16(0xFFFF),
		aarch64codegen.MovShift0,
	)
	assert.Equal(t, expected, uint32(actual))
}

func TestMovk1(t *testing.T) {
	expected := assembleInstruction(t, "MOVK X10, #0, LSL #0x30")
	actual := aarch64codegen.MOVK(
		aarch64codegen.X10,
		aarch64codegen.Immediate16(0),
		aarch64codegen.MovShift48,
	)
	assert.Equal(t, expected, uint32(actual))
}

func TestMovk2(t *testing.T) {
	expected := assembleInstruction(t, "MOVK X29, #0xFFFE, LSL #16")
	actual := aarch64codegen.MOVK(
		aarch64codegen.X29,
		aarch64codegen.Immediate16(0xFFFE),
		aarch64codegen.MovShift16,
	)
	assert.Equal(t, expected, uint32(actual))
}

func TestMovn1(t *testing.T) {
	expected := assembleInstruction(t, "MOVN X2, #0xBEEF, LSL #32")
	actual := aarch64codegen.MOVN(
		aarch64codegen.X2,
		aarch64codegen.Immediate16(0xBEEF),
		aarch64codegen.MovShift32,
	)
	assert.Equal(t, expected, uint32(actual))
}

func TestMovn2(t *testing.T) {
	expected := assembleInstruction(t, "MOVN X21, #0x0, LSL #0")
	actual := aarch64codegen.MOVN(
		aarch64codegen.X21,
		aarch64codegen.Immediate16(0),
		aarch64codegen.MovShift0,
	)
	assert.Equal(t, expected, uint32(actual))
}
