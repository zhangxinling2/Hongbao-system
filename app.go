package Hongbao_system

import (
	_ "github.com/zhangxinling2/account/core/accounts"
	"github.com/zhangxinling2/infra"
	"github.com/zhangxinling2/infra/base"
	"github.com/zhangxinling2/resk/apis/gorpc"
	_ "github.com/zhangxinling2/resk/apis/web"
	_ "github.com/zhangxinling2/resk/core/envelopes"
	"github.com/zhangxinling2/resk/jobs"
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
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}
