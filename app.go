package Hongbao_system

import (
	_ "resk/apis/web"
	_ "resk/core/accounts"
	_ "resk/core/envelopes"
	"resk/infra"
	"resk/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDataBaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisApplicationStarter{})
	infra.Register(&infra.WebStarter{})
}
