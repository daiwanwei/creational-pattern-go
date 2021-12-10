package robots

import "fmt"

type Arm struct {
	Length int
}

func (robot *Arm) Execute() {
	fmt.Printf("length of my arm is %d \n", robot.Length)
}

func (robot *Arm) Clone() Arm {
	return Arm{
		robot.Length,
	}
}
