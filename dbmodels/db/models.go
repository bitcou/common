package adminmodels

import (
	"github.com/bitcou/common/dbmodels/graph/model"
)

/*
These models are for representing the low level arquitechture of our database, GraphQL public API should make tose of
ToGQL() methods to only expose some models fields to the end user.
*/

type Product struct {
	//  Product ID
	ID int `json:"id" gorm:"primary_key"`
	//  Indicates if the product is available
	Available bool `json:"available"`
	//  Product currency
	Currency string `json:"currency"`
	//  Product description
	Description string `json:"description"`
	//  ADMIN Set Product description
	CustomDescription string `json:"customDescription"`
	//  Absolute product discount, expressed in net amount
	DiscountAbsolute float64 `json:"discountAbsolute"`
	//  Product discount percentage, expressed as a decimal from 0 to 1 ***
	DiscountPercentage float64 `json:"discountPercentage"`
	//  Fixed maximum price of the product
	FixedMaxPrice float64 `json:"fixedMaxPrice"`
	//  Fixed minimum price of the product
	FixedMinPrice float64 `json:"fixedMinPrice"`
	//  Product name
	FullName string `json:"fullName"`
	//  ADMIN Set Product name
	CustomFullName string `json:"customFullName"`
	//  Indicates if the product has a discount
	HasDiscount bool `json:"hasDiscount"`
	//  Indicates if the product has a fixed price
	IsFixedPrice bool `json:"isFixedPrice"`
	//  Indicates if the product is premium
	IsPremium bool `json:"isPremium"`
	//  Product country, expressed with ISO 3166 Alpha-2 code
	Locale string `json:"locale"`
	//  Online terms and conditions of the product represented in a string, in some cases with urls in between
	OnlineTc string `json:"onlineTc"`
	//  Original product ID
	OriginalID string `json:"originalID"`
	//  MetaProvider ID
	MetaProviderID int `json:"metaProviderID"`
	//  MetaProvider data
	MetaProvider *model.MetaProvider `json:"metaProvider"`
	//  Provider ID
	ProviderID int `json:"providerID"`
	//  Provider data
	Provider *model.Provider `json:"provider"`
	//  Instructions to redeem the product
	RedeemInstructions string `json:"redeemInstructions"`
	//  Site to redeem the product
	RedeemSite string `json:"redeemSite"`
	//  Whether the product requires user identity
	RequiresUserIdentity bool `json:"requiresUserIdentity"`
	//  Terms and conditions of the product represented in a string
	Tc string `json:"tc"`
	//  URL Image of the product
	URLImage string `json:"urlImage"`
	//  Array containing the countries where the product can be found
	Countries []*model.Country `json:"countries" gorm:"many2many:product_countries;"`
	//  ***
	Variants []*model.Variant `json:"variants"`
	//  Array with categories where the product can be found
	Categories []*model.Category `json:"categories" gorm:"many2many:product_categories;"`
	//  Special field for provider specific product info. Ex store for Baluwo or ean for Ding.
	ProductType string `json:"productType"`
}

func (p *Product) ToGQL() *model.Product {
	// Custom param check
	description := p.Description
	if p.CustomDescription != "" {
		description = p.CustomDescription
	}
	name := p.FullName
	if p.CustomFullName != "" {
		name = p.CustomFullName
	}
	return &model.Product{
		ID:                   p.ID,
		Available:            p.Available,
		Currency:             p.Currency,
		Description:          description,
		DiscountAbsolute:     p.DiscountAbsolute,
		DiscountPercentage:   p.DiscountPercentage,
		FixedMaxPrice:        p.FixedMaxPrice,
		FixedMinPrice:        p.FixedMinPrice,
		FullName:             name,
		HasDiscount:          p.HasDiscount,
		IsFixedPrice:         p.IsFixedPrice,
		IsPremium:            p.IsPremium,
		Locale:               p.Locale,
		OnlineTc:             p.OnlineTc,
		OriginalID:           p.OriginalID,
		MetaProviderID:       p.ProviderID,
		MetaProvider:         p.MetaProvider,
		ProviderID:           p.ProviderID,
		Provider:             p.Provider,
		RedeemInstructions:   p.RedeemInstructions,
		RedeemSite:           p.RedeemSite,
		RequiresUserIdentity: p.RequiresUserIdentity,
		Tc:                   p.Tc,
		URLImage:             p.URLImage,
		Countries:            p.Countries,
		Variants:             p.Variants,
		Categories:           p.Categories,
	}
}
