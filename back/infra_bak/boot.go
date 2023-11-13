package infra

import "github.com/tietang/props/kvs"

type BootApplication struct {
	starterContext StarterContext
	conf           kvs.ConfigSource
}

func New(conf kvs.ConfigSource) *BootApplication {
	b := &BootApplication{
		starterContext: StarterContext{},
		conf:           conf,
	}
	b.starterContext[KeyProps] = conf
	return b
}
func (b *BootApplication) Start() {
	b.init()
	b.setup()
	b.start()
}
func (b *BootApplication) init() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Init(b.starterContext)
	}
}
func (b *BootApplication) setup() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Setup(b.starterContext)
	}
}
func (b *BootApplication) start() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Start(b.starterContext)
		//if starter.StartBlocking() {
		//	if i+1 == len(StarterRegister.AllStarters()) {
		//		starter.Start(b.starterContext)
		//	} else {
		//		go starter.Start(b.starterContext)
		//	}
		//} else {
		//	starter.Start(b.starterContext)
		//}
	}
}
