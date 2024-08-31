package registers

import "fmt"

type GPRegister uint8

const (
	X0 GPRegister = iota
	X1
	X2
	X3
	X4
	X5
	X6
	X7
	X8
	X9
	X10
	X11
	X12
	X13
	X14
	X15
	X16
	X17
	X18
	X19
	X20
	X21
	X22
	X23
	X24
	X25
	X26
	X27
	X28
	X29
	X30
)

func (r GPRegister) Validate() error {
	if r > X30 {
		return fmt.Errorf("invalid general purpose register %d", r)
	}
	return nil
}

func (r GPRegister) Binary() uint32 {
	return uint32(r)
}

func (r GPRegister) String() string {
	return fmt.Sprintf("X%d", r)
}
