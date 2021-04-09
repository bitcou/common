package model

func (c *ClientInput) ToClientModel(id *int) Client {
	newClient := Client{
		BusinessTaxID:   c.BusinessTaxID,
		Name:            c.Name,
		AddressStreet:   c.Address.AddressStreet,
		AddressPc:       c.Address.AddressPc,
		AddressCity:     c.Address.AddressCity,
		AddressState:    c.Address.AddressState,
		AddressCountry:  c.Address.AddressCountry,
		MonthlyFee:      c.MonthlyFee,
		Tc:              c.Tc,
		ContactName:     c.ContactDetails.ContactName,
		ContactLastName: c.ContactDetails.ContactLastName,
		ContactTitle:    c.ContactDetails.ContactTitle,
		ContactEmail:    c.ContactDetails.ContactEmail,
		IsPremium:       c.IsPremium,
		IsAdmin:         false,
		UserName:        c.UserName,
	}
	if id != nil {
		newClient.ID = *id
	}
	return newClient
}
func (c *Client) FromClientModel(input ClientInput) {

	c.BusinessTaxID = input.BusinessTaxID
	c.Name = input.Name
	c.AddressPc = input.Address.AddressPc
	c.AddressCity = input.Address.AddressCity
	c.AddressState = input.Address.AddressState
	c.AddressCountry = input.Address.AddressCountry
	c.AddressStreet = input.Address.AddressStreet
	c.MonthlyFee = input.MonthlyFee
	c.Tc = input.Tc
	c.ContactName = input.ContactDetails.ContactName
	c.ContactLastName = input.ContactDetails.ContactLastName
	c.ContactTitle = input.ContactDetails.ContactTitle
	c.ContactEmail = input.ContactDetails.ContactEmail
	c.IsPremium = input.IsPremium
	c.IsAdmin = false
	c.UserName = input.UserName
}
