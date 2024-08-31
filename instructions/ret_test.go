package instructions_test

import (
	"testing"

	"alon.kr/x/aarch64codegen/instructions"
	"alon.kr/x/aarch64codegen/registers"
)

func TestRet1(t *testing.T) {
	AssertExpectedInstruction(t, "RET", instructions.RET(registers.X30))
}

func TestRet2(t *testing.T) {
	AssertExpectedInstruction(t, "RET X0", instructions.RET(registers.X0))
}

func TestRet3(t *testing.T) {
	AssertExpectedInstruction(t, "RET X29", instructions.RET(registers.X29))
}
