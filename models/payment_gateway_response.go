package models

type PaymentGatewayResponse struct {
	IsSuccess       bool
	Message         string
	HTMLFormContent string
}
