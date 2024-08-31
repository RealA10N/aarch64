package registers

import "fmt"

type GPorSPRegister uint8

const SP GPorSPRegister = 31

func (r GPorSPRegister) Validate() error {
	if r > SP {
		return fmt.Errorf("invalid general purpose or SP register: %d", r)
	}
	return nil
}

func (r GPorSPRegister) Binary() uint32 {
	return uint32(r)
}

func (r GPorSPRegister) String() string {
	if r == SP {
		return "SP"
	}
	return fmt.Sprintf("X%d", r)
}
