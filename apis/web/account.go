package web

import (
	"github.com/kataras/iris/v12"
	"github.com/zhangxinling2/infra"
	"github.com/zhangxinling2/infra/base"
	service "resk/services"
)

const (
	ResCodeBizTransferedFailure = base.ResCode(6010)
)

func init() {
	infra.RegisterApi(new(AccountApi))
}

type AccountApi struct {
	service service.AccountService
}

func (a *AccountApi) Init() {
	a.service = service.GetAccountService()
	groupRouter := base.Iris().Party("/v1/account")
	groupRouter.Post("/create", a.createHandler)
	groupRouter.Post("/transfer", a.transferHandler)
	groupRouter.Get("/get", a.getAccountHandler)
	groupRouter.Get("/envelop/get", a.getEnvelopeAccountHandler)
}
func (a *AccountApi) createHandler(context iris.Context) {
	account := service.AccountCreatedDTO{}
	err := context.ReadJSON(&account)
	r := base.Res{
		Code:    base.ResCodeOk,
		Message: "",
		Date:    nil,
	}
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = err.Error()
		context.JSON(r)
		return
	}

	dto, err := a.service.CreateAccount(account)
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = err.Error()
	}
	r.Date = dto
	context.JSON(&r)
}

func (a *AccountApi) transferHandler(context iris.Context) {
	transfer := service.AccountTransferDTO{}
	err := context.ReadJSON(&transfer)
	r := base.Res{
		Code:    base.ResCodeOk,
		Message: "",
		Date:    nil,
	}
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = err.Error()
		context.JSON(r)
		return
	}
	status, err := a.service.Transfer(transfer)
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = err.Error()
	}
	r.Date = status
	if status != service.TransferSuccess {
		r.Code = ResCodeBizTransferedFailure
		r.Message = err.Error()
	}
	context.JSON(&r)
}
func (a *AccountApi) getEnvelopeAccountHandler(ctx iris.Context) {
	userId := ctx.URLParam("userId")
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if userId == "" {
		r.Code = base.ResCodeRequestParamsError
		r.Message = "ID不能为空"
		ctx.JSON(&r)
		return
	}
	account := a.service.GetEnvelopeAccountByUserId(userId)
	r.Date = account
	ctx.JSON(&r)
}
func (a *AccountApi) getAccountHandler(ctx iris.Context) {
	userId := ctx.URLParam("accountNo")
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if userId == "" {
		r.Code = base.ResCodeRequestParamsError
		r.Message = "账户编号不能为空"
		ctx.JSON(&r)
		return
	}
	account := a.service.GetAccount(userId)
	r.Date = account
	ctx.JSON(&r)
}
