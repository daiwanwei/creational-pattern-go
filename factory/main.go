package main

import (
	"creational-pattern-go/factory/sports"
	"fmt"
)

func main() {
	adidas := sports.NewAdidasFactory()
	nike := sports.NewNikeFactory()
	adidasShoe := adidas.MakeShoe()
	nikeShoe := nike.MakeShoe()
	fmt.Println(adidasShoe.GetSize())
	fmt.Println(nikeShoe.GetSize())
}
