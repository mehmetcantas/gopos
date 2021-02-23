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

type CustomerBuilder struct {
	customer *Customer
}

func NewCustomerBuilder() *CustomerBuilder {
	return &CustomerBuilder{&Customer{}}
}

func (c *CustomerBuilder) NameIs(name string) *CustomerBuilder {
	c.customer.FullName = name
	return c
}
func (c *CustomerBuilder) EmailIs(email string) *CustomerBuilder {
	c.customer.EmailAddress = email
	return c
}
func (c *CustomerBuilder) ShipTo(sAddress string, sZipCode string) *CustomerBuilder {
	c.customer.ShippingAddress = sAddress
	c.customer.ShippingAddressZipCode = sZipCode
	return c
}

func (c *CustomerBuilder) BillTo(bAddress string, bZipCode string) *CustomerBuilder {
	c.customer.BillingAddress = bAddress
	c.customer.BillingAddressZipCode = bZipCode
	return c
}
func (c *CustomerBuilder) IpAddress(ipAddress string) *CustomerBuilder {
	c.customer.IPAddress = ipAddress
	return c
}
func (c *CustomerBuilder) IsCompany(isCompany bool) *CustomerBuilder {
	c.customer.IsCompany = isCompany
	return c
}
func (c *CustomerBuilder) WithID(id string) *CustomerBuilder {
	c.customer.CustomerID = id
	return c
}
