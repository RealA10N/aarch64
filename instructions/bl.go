package instructions

import (
	"alon.kr/x/aarch64codegen/immediates"
)

type Bl uint32

func BL(offset immediates.Offset26Align4) Bl {
	return Bl(0b100101<<26 | offset.Binary())
}

func (i Bl) String() string {
	offset := immediates.Offset26Align4(i & 0x3FFFFFF)
	return "BL " + offset.String()
}

func (i Bl) Binary() uint32 {
	return uint32(i)
}
