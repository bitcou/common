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
