package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"github.com/stretchr/testify/assert"
)

func AssertExpectedBlInstruction(t *testing.T, expected string, offset int32) {
	off26, err := immediates.NewOffset26Align4(offset)
	assert.NoError(t, err)
	instruction := instructions.BL(off26)
	AssertExpectedInstruction(t, expected, instruction)
}

func TestBl0(t *testing.T) {
	AssertExpectedBlInstruction(t, "BL #0", 0)
}

func TestBl1(t *testing.T) {
	AssertExpectedBlInstruction(t, "BL #4", 4)
}

func TestBlMinus1(t *testing.T) {
	AssertExpectedBlInstruction(t, "BL #-4", -4)
}

func TestBlMax(t *testing.T) {
	AssertExpectedBlInstruction(t, "BL #134217724", 0x7FFFFFC)
}

func TestBlMin(t *testing.T) {
	AssertExpectedBlInstruction(t, "BL #-134217728", -0x8000000)
}
