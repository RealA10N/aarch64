package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestLdr1(t *testing.T) {
	AssertExpectedInstruction(t, "LDR X0, [X1], #8", instructions.LDR(
		registers.X0,
		registers.GPorSPRegister(registers.X1),
		immediates.Immediate9(8),
	))
}
