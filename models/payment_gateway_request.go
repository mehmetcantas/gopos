package models

type PaymentGatewayRequest struct {
	CardHolderName   string
	CardNumber       string
	CVV              string
	ExpireMonth      string
	ExpireYear       string
	OrderNumber      string
	OrderTotal       float64
	InstallmentCount int
	CurrencyCode     string
	LanguageCode     string
	CardType         string
	SuccessURL       string
	FailURL          string
	Customer         *Customer
}
