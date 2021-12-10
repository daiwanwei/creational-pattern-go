package sports

type adidasFactory struct {
}

func NewAdidasFactory() Factory {
	return &adidasFactory{}
}

func (factory *adidasFactory) MakeShoe() Shoe {
	return &adidasShoe{
		shoe{
			size: 14,
			logo: "adidas",
		},
	}
}

type adidasShoe struct {
	shoe
}
