package computers

type MobilePhoneBuilder struct {
	Computer
}

func (builder *MobilePhoneBuilder) SetStructure() Builder {
	builder.Structure = "mobile phone"
	return builder
}

func (builder *MobilePhoneBuilder) SetMotherBoard() Builder {
	builder.MotherBoard = 1
	return builder
}

func (builder *MobilePhoneBuilder) SetCPUs() Builder {
	builder.CPU = 1
	return builder
}

func (builder *MobilePhoneBuilder) SetDisk() Builder {
	builder.Disk = 1
	return builder
}

func (builder *MobilePhoneBuilder) SetDRAMs() Builder {
	builder.DRAM = 1
	return builder
}

func (builder *MobilePhoneBuilder) GetComputer() Computer {
	return builder.Computer
}
