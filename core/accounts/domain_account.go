package accounts

import (
	"context"
	"errors"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"resk/infra/base"
	services "resk/services"
)

type accountDomain struct {
	account    Account
	accountLog AccountLog
}

func NewAccountDomain() *accountDomain {
	return new(accountDomain)
}
func (domain *accountDomain) createAccountNo() {
	domain.account.AccountNo = ksuid.New().Next().String()
}
func (domain *accountDomain) createAccountLogNo() {
	domain.accountLog.LogNo = ksuid.New().Next().String()
}
func (domain *accountDomain) createAccountLog() {
	domain.accountLog = AccountLog{}
	domain.createAccountLogNo()
	domain.accountLog.TradeNo = domain.accountLog.LogNo

	domain.accountLog.UserId = domain.account.UserId
	domain.accountLog.AccountNo = domain.account.AccountNo
	domain.accountLog.Username = domain.account.UserName.String

	domain.accountLog.TargetUserId = domain.account.UserId
	domain.accountLog.TargetAccountNo = domain.account.AccountNo
	domain.accountLog.TargetUsername = domain.account.UserName.String

	domain.accountLog.Balance = domain.account.Balance
	domain.accountLog.Amount = domain.account.Balance

	domain.accountLog.ChangeFlag = services.FlagAccountCreated
	domain.accountLog.ChangeType = services.AccountCreated
}
func (domain *accountDomain) Create(dto services.AccountDTO) (*services.AccountDTO, error) {
	domain.account = Account{}
	domain.account.FromDTO(&dto)
	domain.createAccountNo()
	domain.account.UserName.Valid = true

	domain.accountLog = AccountLog{}
	domain.createAccountLog()

	accountDao := AccountDao{}
	accountLogDAO := AccountLogDAO{}
	var rdto *services.AccountDTO
	err := base.Tx(func(runner *dbx.TxRunner) error {
		accountDao.runner = runner
		accountLogDAO.runner = runner
		rs, err := accountDao.Insert(&domain.account)
		if err != nil {
			return err
		}
		if rs <= 0 {
			return errors.New("创建账户失败")
		}
		rs, err = accountLogDAO.Insert(&domain.accountLog)
		if err != nil {
			return err
		}
		if rs <= 0 {
			return errors.New("创建账户流水失败")
		}
		domain.account = *accountDao.GetOne(domain.account.AccountNo)
		return nil
	})
	rdto = domain.account.ToDTO()
	return rdto, err
}
func (domain *accountDomain) Transfer(dto services.AccountTransferDTO) (status services.TransferStatus, err error) {
	err = base.Tx(func(runner *dbx.TxRunner) error {
		ctx := base.WithValueContext(context.Background(), runner)
		status, err = domain.TransferWithContextTx(ctx, dto)
		return err
	})
	return status, err
}
func (domain *accountDomain) TransferWithContextTx(ctx context.Context, dto services.AccountTransferDTO) (status services.TransferStatus, err error) {
	amount := dto.Amount
	if dto.ChangeFlag == services.FlagTransferOut {
		amount = amount.Mul(decimal.NewFromFloat(-1))
	}
	domain.accountLog = AccountLog{}
	domain.accountLog.FromTransferDTO(dto)
	domain.createAccountLogNo()

	err = base.TxContext(ctx, func(runner *dbx.TxRunner) error {
		accountDao := AccountDao{runner: runner}
		accountLogDao := AccountLogDAO{runner: runner}

		rows, err := accountDao.UpdateBalance(dto.TradeBody.AccountNo, amount)
		if err != nil {
			status = services.TransferStatusFailure
			return err
		}
		if rows <= 0 && dto.ChangeFlag == services.FlagTransferOut {
			status = services.TransferSufficientFunds
			return errors.New("余额不足")
		}
		account := accountDao.GetOne(dto.TradeBody.AccountNo)
		if account == nil {
			return errors.New("红包账户不存在")
		}
		domain.account = *account
		domain.accountLog.Balance = domain.account.Balance
		id, err := accountLogDao.Insert(&domain.accountLog)
		if id <= 0 || err != nil {
			status = services.TransferStatusFailure
			return errors.New("账户流水创建失败")
		}
		return nil
	})
	if err != nil {
		logrus.Error(err)
	} else {
		status = services.TransferSuccess
	}
	return status, err
}
func (domain *accountDomain) GetAccount(accountNo string) *services.AccountDTO {
	accountDao := AccountDao{}
	var account *Account
	err := base.Tx(func(runner *dbx.TxRunner) error {
		accountDao.runner = runner
		account = accountDao.GetOne(accountNo)
		return nil
	})
	if err != nil {
		return nil
	}
	if account == nil {
		return nil
	}
	return account.ToDTO()
}
func (domain *accountDomain) GetEnvelopeAccountByUserId(userId string) *services.AccountDTO {
	accountDao := AccountDao{}
	var account *Account
	err := base.Tx(func(runner *dbx.TxRunner) error {
		accountDao.runner = runner
		account = accountDao.GetByUserId(userId, int(services.EnvelopeAccountType))
		return nil
	})
	if err != nil {
		return nil
	}
	if account == nil {
		return nil
	}
	return account.ToDTO()
}
func (domain *accountDomain) GetAccountLog(logNo string) *services.AccountLogDTO {
	accountLogDao := AccountLogDAO{}
	var accountLog *AccountLog
	err := base.Tx(func(runner *dbx.TxRunner) error {
		accountLogDao.runner = runner
		accountLog = accountLogDao.GetOne(logNo)
		return nil
	})
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return accountLog.ToDTO()
}
func (domain *accountDomain) GetAccountLogByTradeNo(tradeNo string) *services.AccountLogDTO {
	accountLogDao := AccountLogDAO{}
	var accountLog *AccountLog
	err := base.Tx(func(runner *dbx.TxRunner) error {
		accountLogDao.runner = runner
		accountLog = accountLogDao.GetByTradeNo(tradeNo)
		return nil
	})
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return accountLog.ToDTO()
}
