package accounts

import (
	"database/sql"
	"github.com/shopspring/decimal"
	services "resk/services"
	"time"
)

type Account struct {
	Id           int64           `db:"id,omitempty"`
	AccountNo    string          `db:"account_no,uni"`
	AccountName  string          `db:"account_name"`
	AccountType  int             `db:"account_type"`
	CurrencyCode string          `db:"currency_code"`
	UserId       string          `db:"user_id"`
	UserName     sql.NullString  `db:"user_name"`
	Balance      decimal.Decimal `db:"balance"`
	Status       int             `db:"status"`
	CreatedAt    time.Time       `db:"created_at,omitempty"`
	UpdatedAt    time.Time       `db:"updated_at,omitempty"`
}

func (a *Account) FromDTO(dto *services.AccountDTO) {
	a.AccountNo = dto.AccountNo
	a.AccountName = dto.AccountName
	a.AccountType = dto.AccountType
	a.CurrencyCode = dto.CurrentCode
	a.UserId = dto.UserId
	a.UserName = sql.NullString{String: dto.UserName, Valid: true}
	a.Balance = dto.Balance
	a.Status = dto.Status
	a.CreatedAt = dto.CreatedAt
	a.UpdatedAt = dto.UpdatedAt
}
func (po *Account) ToDTO() *services.AccountDTO {
	dto := &services.AccountDTO{}
	dto.AccountNo = po.AccountNo
	dto.AccountName = po.AccountName
	dto.AccountType = po.AccountType
	dto.CurrentCode = po.CurrencyCode
	dto.UserId = po.UserId
	dto.UserName = po.UserName.String
	dto.Balance = po.Balance
	dto.Status = po.Status
	dto.CreatedAt = po.CreatedAt
	dto.UpdatedAt = po.UpdatedAt
	return dto
}
