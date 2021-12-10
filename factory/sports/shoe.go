package sports

type Shoe interface {
	SetSize(size int)
	GetSize() int
	SetLogo(logo string)
	getLogo() string
}

type shoe struct {
	size int
	logo string
}

func (s *shoe) SetSize(size int) {
	s.size = size
}

func (s *shoe) GetSize() int {
	return s.size
}

func (s *shoe) SetLogo(logo string) {
	s.logo = logo
}

func (s *shoe) getLogo() string {
	return s.logo
}
