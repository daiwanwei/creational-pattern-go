package robots

import "fmt"

type Human struct {
	Name string
	Arms []Arm
}

func (robot *Human) Execute() {
	fmt.Printf("My name is %s \n", robot.Name)
	for _, arm := range robot.Arms {
		arm.Execute()
	}
}

func (robot *Human) Clone() Human {
	var arms []Arm
	for _, arm := range robot.Arms {
		arms = append(arms, arm.Clone())
	}
	return Human{
		robot.Name,
		arms,
	}
}
