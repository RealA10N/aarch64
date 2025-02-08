package immediates

import "fmt"

// Offset types are signed, in contrast to immediate types which are unsigned.
// The invariant that is maintained is that the lower N bits are always represent
// the offset in it's unsigned form.

type Offset26Align4 int32

func NewOffset26Align4(offset int32) (Offset26Align4, error) {
	if offset%4 != 0 {
		return 0, fmt.Errorf("offset must be a multiple of 4")
	}

	offset /= 4
	if offset < -0x2000000 || offset > 0x1FFFFFF {
		return 0, fmt.Errorf("offset overflows 26 bits")
	}

	return Offset26Align4(offset), nil
}

func (o Offset26Align4) Binary() uint32 {
	return uint32(o) & 0x3FFFFFF
}

func (o Offset26Align4) Signed() int32 {
	return int32(o.Binary()) << 6 >> 6
}

func (o Offset26Align4) String() string {
	return fmt.Sprintf("#%d", o.Signed()*4)
}
