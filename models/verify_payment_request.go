package models

import "net/url"

type VerifyPaymentRequest struct {
	BankName   string
	BankParams url.Values
}
