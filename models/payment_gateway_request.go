package models

import (
	"github.com/mehmetcantas/gopos/models/card_type"
	"github.com/mehmetcantas/gopos/models/currency"
)

type PaymentGatewayRequest struct {
	CardHolderName, CardNumber, CVV, ExpireMonth, ExpireYear string
	OrderNumber                                              string
	OrderTotal                                               float64
	InstallmentCount                                         int
	LanguageCode                                             string
	SuccessURL, FailURL                                      string
	CurrencyCode                                             currency.CurrencyCode
	CardType                                                 card_type.CardType
	Customer                                                 *Customer
}

type PaymentGatewayRequestBuilder struct {
	paymentGatewayRequest *PaymentGatewayRequest
}

func NewPaymentGatewayRequestBuilder() *PaymentGatewayRequestBuilder {
	return &PaymentGatewayRequestBuilder{&PaymentGatewayRequest{}}
}

func (p *PaymentGatewayRequestBuilder) Card(holderName string, cardNumber string, CVV string) *PaymentGatewayRequestBuilder {
	p.paymentGatewayRequest.CardHolderName = holderName
	p.paymentGatewayRequest.CardNumber = cardNumber
	p.paymentGatewayRequest.CVV = CVV
	return p
}

func (p *PaymentGatewayRequestBuilder) ExpireAt(expireMonth string, expireYear string) *PaymentGatewayRequestBuilder {
	p.paymentGatewayRequest.ExpireMonth = expireMonth
	p.paymentGatewayRequest.ExpireYear = expireYear
	return p
}
func (p *PaymentGatewayRequestBuilder) Type(cardType card_type.CardType) *PaymentGatewayRequestBuilder {
	p.paymentGatewayRequest.CardType = cardType
	return p
}

func (p *PaymentGatewayRequestBuilder) Currency(code currency.CurrencyCode) *PaymentGatewayRequestBuilder {
	p.paymentGatewayRequest.CurrencyCode = code
	return p
}
func (p *PaymentGatewayRequestBuilder) Language(lCode string) *PaymentGatewayRequestBuilder {
	p.paymentGatewayRequest.LanguageCode = lCode
	return p
}
func (p *PaymentGatewayRequestBuilder) ForOrder(oNumber string, oTotal float64) *PaymentGatewayRequestBuilder {

	p.paymentGatewayRequest.OrderNumber = oNumber
	p.paymentGatewayRequest.OrderTotal = oTotal
	return p
}
func (p *PaymentGatewayRequestBuilder) WithInstallment(installment int) *PaymentGatewayRequestBuilder {
	p.paymentGatewayRequest.InstallmentCount = installment
	return p
}
func (p *PaymentGatewayRequestBuilder) ToCustomer(customer *Customer) *PaymentGatewayRequestBuilder {
	p.paymentGatewayRequest.Customer = customer
	return p
}
func (p *PaymentGatewayRequestBuilder) InSuccessReturns(successUrl string) *PaymentGatewayRequestBuilder {
	p.paymentGatewayRequest.SuccessURL = successUrl
	return p
}

func (p *PaymentGatewayRequestBuilder) InFailReturns(failUrl string) *PaymentGatewayRequestBuilder {
	p.paymentGatewayRequest.FailURL = failUrl
	return p
}

func (p *PaymentGatewayRequestBuilder) Build() *PaymentGatewayRequest {
	return p.paymentGatewayRequest
}
