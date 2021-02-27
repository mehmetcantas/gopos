package models

type Customer struct {
	EmailAddress           string
	ShippingCustomerName   string
	ShippingAddress        string
	ShippingAddressZipCode string
	BillingCustomerName    string
	BillingAddress         string
	BillingAddressZipCode  string
	IPAddress              string
	IsBillingToCompany     bool
	CustomerID             string
}

type CustomerBuilder struct {
	customer *Customer
}

func NewCustomerBuilder() *CustomerBuilder {
	return &CustomerBuilder{&Customer{}}
}

func (c *CustomerBuilder) EmailIs(email string) *CustomerBuilder {
	c.customer.EmailAddress = email
	return c
}
func (c *CustomerBuilder) ShipTo(sName string, sAddress string, sZipCode string) *CustomerBuilder {
	c.customer.ShippingAddress = sAddress
	c.customer.ShippingCustomerName = sName
	c.customer.ShippingAddressZipCode = sZipCode
	return c
}

func (c *CustomerBuilder) BillTo(bName string, bAddress string, bZipCode string) *CustomerBuilder {
	c.customer.BillingAddress = bAddress
	c.customer.BillingCustomerName = bName
	c.customer.BillingAddressZipCode = bZipCode
	return c
}
func (c *CustomerBuilder) IpAddress(ipAddress string) *CustomerBuilder {
	c.customer.IPAddress = ipAddress
	return c
}
func (c *CustomerBuilder) IsBillingToCompany(isCompany bool) *CustomerBuilder {
	c.customer.IsBillingToCompany = isCompany
	return c
}
func (c *CustomerBuilder) WithID(id string) *CustomerBuilder {
	c.customer.CustomerID = id
	return c
}

func (c *CustomerBuilder) Build() *Customer {
	return c.customer
}
