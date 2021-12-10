package computers

type Builder interface {
	SetStructure() Builder
	SetMotherBoard() Builder
	SetCPUs() Builder
	SetDisk() Builder
	SetDRAMs() Builder
	GetComputer() Computer
}

type Computer struct {
	Structure   string
	MotherBoard int
	CPU         int
	Disk        int
	DRAM        int
}

type ManufacturingDirector struct {
	builder Builder
}

func (director *ManufacturingDirector) SetBuilder(builder Builder) {
	director.builder = builder
}

func (director *ManufacturingDirector) Manufacture() {
	director.builder.
		SetStructure().
		SetMotherBoard().
		SetCPUs().
		SetDRAMs().
		SetDisk()
}
