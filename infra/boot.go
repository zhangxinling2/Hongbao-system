package infra

import "github.com/tietang/props/kvs"

type BootApplication struct {
	starterContext StarterContext
	conf           kvs.ConfigSource
}

func New(conf kvs.ConfigSource) *BootApplication {
	return &BootApplication{
		starterContext: StarterContext{},
		conf:           conf,
	}
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
	for i, starter := range StarterRegister.AllStarters() {

		if starter.StartBlocking() {
			if i+1 == len(StarterRegister.AllStarters()) {
				starter.Start(b.starterContext)
			} else {
				go starter.Start(b.starterContext)
			}
		} else {
			starter.Start(b.starterContext)
		}
	}
}
