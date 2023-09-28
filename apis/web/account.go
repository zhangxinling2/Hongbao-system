package web

import (
	"github.com/kataras/iris/v12"
	"resk/infra"
	"resk/infra/base"
	service "resk/services"
)

const (
	ResCodeBizTransferedFailure = base.ResCode(6010)
)

func init() {
	infra.RegisterApi(new(AccountApi))
}

type AccountApi struct {
}

func (a *AccountApi) Init() {
	groupRouter := base.Iris().Party("/v1/account")
	groupRouter.Post("/create", createHandler)
	groupRouter.Post("/transfer", transferHandler)
	groupRouter.Get("/get", getAccountHandler)
	groupRouter.Get("/envelop/get", getEnvelopeAccountHandler)
}
func createHandler(context iris.Context) {
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
	ser := service.GetAccountService()
	dto, err := ser.CreateAccount(account)
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = err.Error()
	}
	r.Date = dto
	context.JSON(&r)
}

func transferHandler(context iris.Context) {
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
	ser := service.GetAccountService()
	status, err := ser.Transfer(transfer)
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
func getEnvelopeAccountHandler(ctx iris.Context) {
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
	ser := service.GetAccountService()
	account := ser.GetEnvelopeAccountByUserId(userId)
	r.Date = account
	ctx.JSON(&r)
}
func getAccountHandler(ctx iris.Context) {
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
	ser := service.GetAccountService()
	account := ser.GetAccount(userId)
	r.Date = account
	ctx.JSON(&r)
}
