package Hongbao_system

import (
	"resk/infra"
	"resk/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDataBaseStarter{})
}
