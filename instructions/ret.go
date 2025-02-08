package instructions

import (
	"alon.kr/x/aarch64codegen/registers"
)

type Ret uint32

func RET(Xn registers.GPRegister) Ret {
	return Ret(0xD65F0000 | (Xn.Binary() << 5))
}

func (r Ret) String() string {
	register := registers.GPRegister((r >> 5) & 0b11111)
	s := "RET"
	if register != registers.X30 {
		s += " " + register.String()
	}
	return s
}

func (r Ret) Binary() uint32 {
	return uint32(r)
}
