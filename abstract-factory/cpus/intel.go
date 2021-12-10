package cpus

type IntelCPU struct {
	cpu
}

type IntelFactory struct {
}

func (factory *IntelFactory) ManufactureCPU() CPU {
	return &IntelCPU{
		cpu{name: "i5", brand: "Intel"},
	}
}
