package example

import (
	"fmt"
	"github.com/tietang/props/ini"
)

func main() {
	conf := ini.NewIniFileConfigSource("config.ini")
	port := conf.GetIntDefault("app.service.port", 18080)
	fmt.Println(port)
}
