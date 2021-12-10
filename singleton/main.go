package main

func main() {
	manager, err := GetManager()
	if err != nil {
		panic(err)
	}
	manager.SayHi()
	manager.SayBye()
	manager2, err := GetManager()
	if err != nil {
		panic(err)
	}
	manager2.SayHi()
	manager2.SayBye()
}
