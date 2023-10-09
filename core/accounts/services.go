package accounts

import (
	"errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"resk/infra/base"
	services "resk/services"
	"sync"
)

var once sync.Once
var _ services.AccountService = new(AccountService)

func init() {
	once.Do(func() {
		services.IAccountService = new(AccountService)
	})
}

type AccountService struct {
}

func (a *AccountService) CreateAccount(dto services.AccountCreatedDTO) (*services.AccountDTO, error) {
	domain := accountDomain{}
	err := base.ValidateStruct(&dto)
	amount, err := decimal.NewFromString(dto.Amount)
	if err != nil {
		return nil, err
	}
	account := services.AccountDTO{
		AccountCreatedDTO: services.AccountCreatedDTO{
			UserId:      dto.UserId,
			UserName:    dto.UserName,
			AccountType: dto.AccountType,
			AccountName: dto.AccountName,
			CurrentCode: dto.CurrentCode,
		},
		Balance: amount,
		Status:  1,
	}
	rdto, err := domain.Create(account)
	return rdto, err
}

func (a *AccountService) Transfer(dto services.AccountTransferDTO) (services.TransferStatus, error) {
	domain := accountDomain{}
	err := base.Validate().Struct(&dto)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			logrus.Error(err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {
				logrus.Error(e.Translate(base.Translator()))
			}
		}
		return services.TransferStatusFailure, err
	}
	amount, err := decimal.NewFromString(dto.AmountStr)
	if err != nil {
		return services.TransferStatusFailure, err
	}
	dto.Amount = amount
	if dto.ChangeFlag == services.FlagTransferOut {
		if dto.ChangeType > 0 {
			return services.TransferStatusFailure, errors.New("转出时flag必须小于0")
		}
	} else {
		if dto.ChangeType < 0 {
			return services.TransferStatusFailure, errors.New("转入时flag必须大于0")
		}
	}
	return domain.Transfer(dto)
}

func (a *AccountService) StoreValue(dto services.AccountTransferDTO) (services.TransferStatus, error) {
	dto.TradeTarget = dto.TradeBody
	dto.ChangeType = services.AccountStoreValue
	dto.ChangeType = services.EnvelopeIncoming
	return a.Transfer(dto)
}

func (a *AccountService) GetEnvelopeAccountByUserId(UserId string) *services.AccountDTO {
	domain := accountDomain{}
	return domain.GetEnvelopeAccountByUserId(UserId)
}

func (a *AccountService) GetAccount(accountNo string) *services.AccountDTO {
	domain := accountDomain{}
	return domain.GetAccount(accountNo)
}
