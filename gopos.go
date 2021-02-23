package gopos

import "github.com/mehmetcantas/gopos/models"

type IPaymentProvider interface {
	PreparePaymentGatewayForm(r *models.PaymentGatewayRequest) (models.PaymentGatewayResponse, error)
	VerifyPayment(r *models.VerifyPaymentRequest) (models.VerifyPaymentResponse, error)
}
