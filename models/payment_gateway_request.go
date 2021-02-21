package models

type PaymentGatewayRequest struct {
	CardHolderName       string
	CardNumber           string
	CVV                  string
	ExpireMonth          string
	ExpireYear           string
	CustomerEmailAddress string
	CompanyName          string
	OrderNumber          string
	OrderTotal           float64
	InstallmentCount     int
	CurrencyCode         string
	LanguageCode         string
	UserID               string
	CustomerIPAddress    string
	CardType             string
	SuccessURL           string
	FailURL              string
}
