package aarch64codegen_test

import (
	"testing"

	"alon.kr/x/aarch64codegen"
	"github.com/stretchr/testify/assert"
)

func TestMovz(t *testing.T) {
	expected := assembleInstruction(t, "MOVZ x0, #0x1234, LSL #32")
	actual := aarch64codegen.MOVZ(
		aarch64codegen.X0,
		aarch64codegen.Immediate16(0x1234),
		aarch64codegen.MovShift32,
	)
	assert.Equal(t, expected, uint32(actual))
}
