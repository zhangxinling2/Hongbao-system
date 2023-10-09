package base

import (
	"fmt"
	"github.com/tietang/props/kvs"
	"resk/infra"
	"sync"
)

var props kvs.ConfigSource

// Props 对外暴露配置实例
func Props() kvs.ConfigSource {
	return props
}

// PropsStarter BaseStarter的实现
type PropsStarter struct {
	infra.BaseStarter
}

func (p *PropsStarter) Init(ctx infra.StarterContext) {
	props = ctx.Props()
	fmt.Println("配置初始化")
	GetSystemAccount()
}

type SystemAccount struct {
	UserId      string
	UserName    string
	AccountName string
	AccountNo   string
}

var systemAccount *SystemAccount
var once sync.Once

func GetSystemAccount() *SystemAccount {
	once.Do(func() {
		systemAccount = new(SystemAccount)
		err := kvs.Unmarshal(Props(), systemAccount, "system.account")
		if err != nil {
			panic(err)
		}
	})
	return systemAccount
}
func GetEnvelopeLink() string {
	link := Props().GetDefault("envelope.link", "/v1/envelope/link")
	return link
}
func GetEnvelopeDomain() string {
	domain := Props().GetDefault("envelope.domain", "http://localhost")
	return domain
}
