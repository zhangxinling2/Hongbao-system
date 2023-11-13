package infra

type Initializer interface {
	Init()
}
type InitializeRegister struct {
	Initializers []Initializer
}

func (i *InitializeRegister) Register(ini Initializer) {
	i.Initializers = append(i.Initializers, ini)
}
