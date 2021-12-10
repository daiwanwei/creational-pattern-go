package sports

type nikeFactory struct {
}

func NewNikeFactory() Factory {
	return &nikeFactory{}
}

func (factory *nikeFactory) MakeShoe() Shoe {
	return &nikeShoe{
		shoe{
			size: 14,
			logo: "nike",
		},
	}
}

type nikeShoe struct {
	shoe
}
