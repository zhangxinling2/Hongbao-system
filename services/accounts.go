package service

import (
	"github.com/shopspring/decimal"
	"github.com/zhangxinling2/infra/base"
	"time"
)

var IAccountService AccountService

func GetAccountService() AccountService {
	base.Check(IAccountService)
	return IAccountService
}

type AccountService interface {
	CreateAccount(dto AccountCreatedDTO) (*AccountDTO, error)
	Transfer(dto AccountTransferDTO) (TransferStatus, error)
	StoreValue(dto AccountTransferDTO) (TransferStatus, error)
	GetEnvelopeAccountByUserId(UserId string) *AccountDTO
	GetAccount(accountNo string) *AccountDTO
}
type AccountCreatedDTO struct {
	UserId      string
	UserName    string
	AccountName string
	AccountType int
	CurrentCode string
	Amount      string
}
type AccountDTO struct {
	AccountCreatedDTO
	AccountNo string
	CreatedAt time.Time
	UpdatedAt time.Time
	Balance   decimal.Decimal
	Status    int
}
type TradeParticipator struct {
	AccountNo string
	UserId    string
	UserName  string
}
type AccountTransferDTO struct {
	TradeNo     string
	TradeBody   TradeParticipator
	TradeTarget TradeParticipator
	AmountStr   string
	Amount      decimal.Decimal
	ChangeType  ChangeType
	ChangeFlag  ChangeFlag
	Decs        string
}
type AccountLogDTO struct {
	LogNo           string          //流水编号 全局不重复字符或数字，唯一性标识
	TradeNo         string          //交易单号 全局不重复字符或数字，唯一性标识
	AccountNo       string          //账户编号 账户ID
	TargetAccountNo string          //账户编号 账户ID
	UserId          string          //用户编号
	UserName        string          //用户名称
	TargetUserId    string          //目标用户编号
	TargetUserName  string          //目标用户名称
	Amount          decimal.Decimal //交易金额,该交易涉及的金额
	Balance         decimal.Decimal //交易后余额,该交易后的余额
	ChangeType      ChangeType      //流水交易类型，0 创建账户，>0 为收入类型，<0 为支出类型，自定义
	ChangeFlag      ChangeFlag      //交易变化标识：-1 出账 1为进账，枚举
	Status          int             //交易状态：
	Decs            string          //交易描述
	CreatedAt       time.Time       //创建时间
}
