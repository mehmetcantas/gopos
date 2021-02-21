package models

type Customer struct {
	FullName               string
	EmailAddress           string
	ShippingAddress        string
	ShippingAddressZipCode string
	BillingAddress         string
	BillingAddressZipCode  string
	IPAddress              string
	IsCompany              bool
	CustomerID             string
}
