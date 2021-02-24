package models

import (
	"github.com/mehmetcantas/gopos/models/card_type"
	"github.com/mehmetcantas/gopos/models/currency"
)

type PaymentRequest struct {
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

type PaymentRequestBuilder struct {
	PaymentRequest *PaymentRequest
}

func NewPaymentRequestBuilder() *PaymentRequestBuilder {
	return &PaymentRequestBuilder{&PaymentRequest{}}
}

func (p *PaymentRequestBuilder) Card(holderName string, cardNumber string, CVV string) *PaymentRequestBuilder {
	p.PaymentRequest.CardHolderName = holderName
	p.PaymentRequest.CardNumber = cardNumber
	p.PaymentRequest.CVV = CVV
	return p
}

func (p *PaymentRequestBuilder) ExpireAt(expireMonth string, expireYear string) *PaymentRequestBuilder {
	p.PaymentRequest.ExpireMonth = expireMonth
	p.PaymentRequest.ExpireYear = expireYear
	return p
}
func (p *PaymentRequestBuilder) Type(cardType card_type.CardType) *PaymentRequestBuilder {
	p.PaymentRequest.CardType = cardType
	return p
}

func (p *PaymentRequestBuilder) Currency(code currency.CurrencyCode) *PaymentRequestBuilder {
	p.PaymentRequest.CurrencyCode = code
	return p
}
func (p *PaymentRequestBuilder) Language(lCode string) *PaymentRequestBuilder {
	p.PaymentRequest.LanguageCode = lCode
	return p
}
func (p *PaymentRequestBuilder) ForOrder(oNumber string, oTotal float64) *PaymentRequestBuilder {

	p.PaymentRequest.OrderNumber = oNumber
	p.PaymentRequest.OrderTotal = oTotal
	return p
}
func (p *PaymentRequestBuilder) WithInstallment(installment int) *PaymentRequestBuilder {
	p.PaymentRequest.InstallmentCount = installment
	return p
}
func (p *PaymentRequestBuilder) ToCustomer(customer *Customer) *PaymentRequestBuilder {
	p.PaymentRequest.Customer = customer
	return p
}
func (p *PaymentRequestBuilder) InSuccessReturns(successUrl string) *PaymentRequestBuilder {
	p.PaymentRequest.SuccessURL = successUrl
	return p
}

func (p *PaymentRequestBuilder) InFailReturns(failUrl string) *PaymentRequestBuilder {
	p.PaymentRequest.FailURL = failUrl
	return p
}

func (p *PaymentRequestBuilder) Build() *PaymentRequest {
	return p.PaymentRequest
}
