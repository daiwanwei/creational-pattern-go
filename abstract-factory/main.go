package main

import "fmt"

func main() {
	factory := NewMacFactory("AMD")
	macbook := factory.ManufactureMacBook()
	fmt.Println(macbook.GetCPUBrand())
}
