package instructions

import (
	"alon.kr/x/aarch64codegen/immediates"
)

type bl uint32

func BL(offset immediates.Offset26Align4) bl {
	return bl(0b100101<<26 | offset.Binary())
}

func (i bl) String() string {
	offset := immediates.Offset26Align4(i & 0x3FFFFFF)
	return "BL " + offset.String()
}

func (i bl) Binary() uint32 {
	return uint32(i)
}
