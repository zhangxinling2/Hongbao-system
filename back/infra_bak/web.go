package infra

var apiInitializerRegister *InitializeRegister = new(InitializeRegister)

func RegisterApi(ai Initializer) {
	apiInitializerRegister.Register(ai)
}
func GetApiInitializers() []Initializer {
	return apiInitializerRegister.Initializers
}

type WebStarter struct {
	BaseStarter
}

func (w *WebStarter) Setup(ctx StarterContext) {
	for _, reg := range GetApiInitializers() {
		reg.Init()
	}
}
