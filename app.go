package Hongbao_system

import (
	"resk/apis/gorpc"
	_ "resk/apis/web"
	_ "resk/core/accounts"
	_ "resk/core/envelopes"
	"resk/infra"
	"resk/infra/base"
	"resk/jobs"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDataBaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.GoRPCStarter{})
	infra.Register(&gorpc.GoRpcApiStarter{})
	infra.Register(&jobs.RefundExpiredJobStarter{})
	infra.Register(&base.IrisApplicationStarter{})
	infra.Register(&infra.WebStarter{})
	infra.Register(&base.HookStarter{})
}
