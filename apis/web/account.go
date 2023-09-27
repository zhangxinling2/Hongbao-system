package web

import (
	"github.com/kataras/iris/v12"
	"resk/infra"
	"resk/infra/base"
	service "resk/services"
)

func init() {
	infra.RegisterApi(new(AccountApi))
}

type AccountApi struct {
}

func (a *AccountApi) Init() {
	groupRouter := base.Iris().Party("/v1/account")
	create(groupRouter)
}
func create(groupRouter iris.Party) {
	groupRouter.Post("/create", func(context iris.Context) {
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
	})
}
