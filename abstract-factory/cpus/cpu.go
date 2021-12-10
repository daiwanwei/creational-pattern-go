package cpus

type CPUFactory interface {
	ManufactureCPU() CPU
}

type CPU interface {
	GetBrand() string
	GetName() string
}

type cpu struct {
	name  string
	brand string
}

func (c *cpu) GetName() string {
	return c.name
}

func (c *cpu) GetBrand() string {
	return c.brand
}
