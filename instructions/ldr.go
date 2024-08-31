package instructions

import (
	"fmt"

	"alon.kr/x/aarch64codegen/immediates"
	"alon.kr/x/aarch64codegen/registers"
)

type ldr uint32

func LDR(
	Xt registers.GPRegister,
	Xn registers.GPorSPRegister,
	imm immediates.Immediate9,
) ldr {
	return ldr(
		0xF8400400 |
			(imm.Binary() << 12) |
			(Xn.Binary() << 5) |
			(Xt.Binary()),
	)
}

func (i ldr) Binary() uint32 {
	return uint32(i)
}

func (i ldr) String() string {
	imm := immediates.Immediate9((i >> 12) & 0x1FF)
	Xn := registers.GPorSPRegister((i >> 5) & 0x1F)
	Xt := registers.GPRegister(i & 0x1F)
	s := fmt.Sprintf("LDR %s, [%s]", Xt, Xn)
	if imm != 0 {
		s += ", " + imm.String()
	}
	return s
}
