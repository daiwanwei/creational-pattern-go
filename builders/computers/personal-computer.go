package computers

type PersonalComputerBuilder struct {
	Computer
}

func (builder *PersonalComputerBuilder) SetStructure() Builder {
	builder.Structure = "personal computer"
	return builder
}

func (builder *PersonalComputerBuilder) SetMotherBoard() Builder {
	builder.MotherBoard = 1
	return builder
}

func (builder *PersonalComputerBuilder) SetCPUs() Builder {
	builder.CPU = 2
	return builder
}

func (builder *PersonalComputerBuilder) SetDisk() Builder {
	builder.Disk = 2
	return builder
}

func (builder *PersonalComputerBuilder) SetDRAMs() Builder {
	builder.DRAM = 4
	return builder
}

func (builder *PersonalComputerBuilder) GetComputer() Computer {
	return builder.Computer
}
