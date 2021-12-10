package main

import "creational-pattern-go/prototype/robots"

func main() {
	arm1 := robots.Arm{
		Length: 10,
	}
	arm2 := robots.Arm{
		Length: 20,
	}
	human := robots.Human{
		Name: "ann",
		Arms: []robots.Arm{
			arm1, arm2,
		},
	}
	human.Execute()
	human2 := human.Clone()
	human2.Execute()
}
