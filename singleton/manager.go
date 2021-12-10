package main

import "fmt"

var (
	managerInstance Manager
)

func GetManager() (instance Manager, err error) {
	if managerInstance == nil {
		instance, err = newManager("kevin")
		if err != nil {
			return nil, err
		}
		managerInstance = instance
		fmt.Println("new manager")
	}
	return managerInstance, nil
}

type Manager interface {
	SayHi()
	SayBye()
}

type manager struct {
	name string
}

func newManager(name string) (instance Manager, err error) {
	return &manager{name: name}, nil
}

func (m *manager) SayHi() {
	fmt.Println(m.name + "say hi")
}

func (m *manager) SayBye() {
	fmt.Println(m.name + "say bye")
}
