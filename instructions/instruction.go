package instructions

import "fmt"

type Instruction interface {
	fmt.Stringer
	Binary() uint32
}
