package cpus

type AmdCPU struct {
	cpu
}

type AMDFactory struct {
}

func (factory *AMDFactory) ManufactureCPU() CPU {
	return &AmdCPU{
		cpu{name: "zen", brand: "AMD"},
	}
}
