package instructions

import (
	"alon.kr/x/aarch64codegen/registers"
)

type ret uint32

func RET(Xn registers.GPRegister) ret {
	return ret(0xD65F0000 | (Xn.Binary() << 5))
}

func (r ret) String() string {
	register := registers.GPRegister((r >> 5) & 0b11111)
	s := "RET"
	if register != registers.X30 {
		s += " " + register.String()
	}
	return s
}

func (r ret) Binary() uint32 {
	return uint32(r)
}
