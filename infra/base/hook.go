package base

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"reflect"
	"resk/infra"
	"syscall"
)

var callback []func()

func Register(cb func()) {
	callback = append(callback, cb)
}

type HookStarter struct {
	infra.BaseStarter
}

func (h *HookStarter) Init(ctx infra.StarterContext) {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		for {
			c := <-sigs
			log.Info("notify:", c)
			for _, cb := range callback {
				cb()
			}
			break
		}
		os.Exit(0)
	}()

}
func (h *HookStarter) Start(ctx infra.StarterContext) {
	starters := infra.GetStarters()
	for _, starter := range starters {
		Register(func() {
			typ := reflect.TypeOf(starter)
			log.Info("[Register Notify Stop]:", typ)
			starter.Stop(ctx)
		})
	}
}
