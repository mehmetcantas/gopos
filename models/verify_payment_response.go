package models

type VerifyPaymentResponse struct {
	IsSuccess      bool
	BankMessage    string
	BankStatusCode string
	BankErrorCode  string
	TransactionID  string
	OrderID        string
	PaidAmount     float64
}
