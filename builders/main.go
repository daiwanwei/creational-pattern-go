package main

import (
	"creational-pattern-go/builders/computers"
	"fmt"
)

func main() {
	director := computers.ManufacturingDirector{}
	phoneBuilder := computers.MobilePhoneBuilder{}
	pcBuilder := computers.PersonalComputerBuilder{}
	director.SetBuilder(&phoneBuilder)
	director.Manufacture()
	fmt.Println(phoneBuilder.Structure)
	director.SetBuilder(&pcBuilder)
	director.Manufacture()
	fmt.Println(pcBuilder.Structure)
}
