package main

import "creational-pattern-go/abstract-factory/cpus"

type MacFactory struct {
	cpuFactory cpus.CPUFactory
}

func NewMacFactory(cpuBrand string) *MacFactory {
	var factory cpus.CPUFactory
	switch cpuBrand {
	case "AMD":
		factory = &cpus.AMDFactory{}
	case "INTEL":
	default:
		factory = &cpus.IntelFactory{}
	}
	return &MacFactory{
		factory,
	}
}

func (factory *MacFactory) ManufactureMacBook() *MacBook {
	cpu := factory.cpuFactory.ManufactureCPU()
	return &MacBook{
		cpu: cpu,
	}
}

type MacBook struct {
	cpu cpus.CPU
}

func (m *MacBook) GetCPUBrand() string {
	return m.cpu.GetBrand()
}
