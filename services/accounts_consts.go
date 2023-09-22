package service

type TransferStatus int8

const (
	TransferStatusFailure   TransferStatus = -1
	TransferSufficientFunds TransferStatus = 0
	TransferSuccess         TransferStatus = 1
)

type ChangeType int8

const (
	AccountCreated    ChangeType = 0
	AccountStoreValue ChangeType = 1
	EnvelopeOutgoing  ChangeType = -2
	EnvelopeIncoming  ChangeType = 2
	EnvelopeExpire    ChangeType = 3
)

type ChangeFlag int8

const (
	FlagAccountCreated ChangeFlag = 0
	FlagTransferOut    ChangeFlag = -1
	FlagTransferIn     ChangeFlag = 1
)

type AccountType int8

const (
	EnvelopeAccountType       AccountType = 1
	SystemEnvelopeAccountType AccountType = 2
)
