//go:generate go run github.com/99designs/gqlgen
package graph

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/bitcou/common/dbmodels/graph/model"
	"github.com/davegardnerisme/phonegeocode"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *gorm.DB
}

func (r *mutationResolver) UpdateClientResolver(id *int, client model.ClientInput) (*model.Client, error) {
	var c model.Client
	query := r.Resolver.DB
	clientData := client.ToClientModel(id)
	if id == nil {
		// todo add API Keyss
		result := query.Create(&clientData)
		if result.Error != nil {
			return &c, result.Error
		}
		c = clientData
	} else {
		query.First(&c, *id)
		c.FromClientModel(client) // Updated Model
		update := r.Resolver.DB
		update = update.Model(&c).Updates(c)
		if update.Error != nil {
			log.Println("error updating client info ", clientData)
			return &c, update.Error
		}
	}
	return &c, nil
}

func (r *mutationResolver) UpdateProductResolver(id int, product model.ProductInput) (*model.ProductAdmin, error) {
	var products []*model.ProductAdmin
	var p *model.ProductAdmin
	query := r.Resolver.DB

	query = query.Where("id = ?", id).Find(&products)

	if query.Error != nil {
		log.Println("error retrieving")
		return p, query.Error
	}

	p = products[0]
	var mapChanges = make(map[string]interface{})

	if product.CustomDescription != "" {
		mapChanges["custom_description"] = product.CustomDescription
	}
	if product.CustomFullName != "" {
		mapChanges["custom_full_name"] = product.CustomFullName
	}

	if product.CustomURLImage != "" {
		mapChanges["custom_url_image"] = product.CustomURLImage
	}

	if product.CustomDiscount != 0 {
		mapChanges["custom_discount"] = product.CustomDiscount
	}

	update := r.Resolver.DB

	update = update.Model(&p).Updates(mapChanges)

	if update.Error != nil {
		log.Println("error updating product ", product)
		return p, update.Error
	}
	return p, nil
}

func (r *queryResolver) PurchasesResolver(filter *model.PurchaseFilter, limit *int, offset *int) ([]*model.Purchase, error) {
	var purchases []*model.Purchase
	query := r.Resolver.DB

	if filter != nil {
		if filter.ID != nil {
			query = query.Where("id = ?", *filter.ID)
		}
		if filter.ClientID != nil && *filter.ClientID > 0 {
			query = query.Where("client_id = ?", *filter.ClientID)
		}
		if filter.ProductID != nil && *filter.ProductID > 0 {
			query = query.Where("product_id = ?", *filter.ProductID)
		}
		if filter.DateRange != nil {
			if filter.DateRange.Start > filter.DateRange.End || filter.DateRange.End < filter.DateRange.Start {
				return purchases, errors.New("Check the start and end values for the date range.")
			} else {
				query = query.Where("timestamp BETWEEN ? AND ?", filter.DateRange.Start, filter.DateRange.End)
			}
		}
		if filter.PriceRange != nil {
			var minPrice = 0.0
			var maxPrice float64
			row := query.Table("purchases").Select("max(total_euro)").Row()
			err := row.Scan(&maxPrice)
			if err != nil {
				log.Println(err)
				return purchases, err
			}

			if filter.PriceRange.MinPrice != nil && *filter.PriceRange.MinPrice > 0 {
				minPrice = *filter.PriceRange.MinPrice
			}
			if filter.PriceRange.MaxPrice != nil && *filter.PriceRange.MaxPrice > 0 {
				maxPrice = *filter.PriceRange.MaxPrice
			}
			if minPrice > maxPrice || maxPrice < minPrice {
				return purchases, errors.New("Check the minimum and maximum values for the price range.")
			} else {
				query = query.Where("total_euro BETWEEN ? AND ?", minPrice, maxPrice)
			}
		}
	}
	if *limit > 0 {
		query = query.Limit(*limit)
	}
	if *offset > 0 {
		query = query.Offset(*offset)
	}

	err := query.Order("id").Preload("Product").Preload("Client").Find(&purchases).Error
	if err != nil {
		return purchases, err
	}
	return purchases, nil
}

func (r *queryResolver) ClientsResolver(filter *model.ClientFilter, limit *int, offset *int, isAdmin bool) ([]*model.Client, error) {
	var clients []*model.Client
	query := r.Resolver.DB

	if filter != nil {
		if !isAdmin {
			if filter.ID != nil {
				query = query.Where("id = ?", *filter.ID)
			}
		}

		if filter.Name != nil && *filter.Name != "" {
			query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *filter.Name))
		}
		if filter.AddressPc != nil {
			query = query.Where("address_pc = ?", *filter.AddressPc)
		}
		if filter.AddressState != nil {
			query = query.Where("address_state = ?", *filter.AddressState)
		}
		if filter.AddressCountry != nil {
			query = query.Where("address_country = ?", *filter.AddressCountry)
		}
	}
	if *limit > 0 {
		query = query.Limit(*limit)
	}
	if *offset > 0 {
		query = query.Offset(*offset)
	}

	err := query.Order("id").Find(&clients).Error
	if err != nil {
		return clients, err
	}
	return clients, nil
}

func (r *queryResolver) ClientInfoResolver(username string, password string) (string, error) {
	// TODO Make this more efficient using GORM
	var clients []*model.Client
	// var apiKeys []model.APIKey
	query := r.Resolver.DB
	// query2 := r.Resolver.DB

	query = query.Where("user_name = ? AND password = ?", username, password).Preload("APIKeys").Find(&clients)
	if query.Error != nil {
		return "", query.Error
	}
	if len(clients) < 0 {
		return "", errors.New("not found")
	}
	if len(clients[0].APIKeys) < 0 {
		return "", errors.New("not found")
	}
	// query2 = query2.Where("client_id = ?", clients[0].ID).Find(&apiKeys)
	// if query2.Error != nil {
	// 	return "", query.Error
	// }
	// if len(apiKeys) < 0 {
	// 	return "", errors.New("not found")
	// }

	return clients[0].APIKeys[0].Key, nil
}

func (r *queryResolver) ProductsResolver(filter *model.ProductFilter, limit *int, offset *int) ([]*model.Product, error) {
	var products []*model.Product
	query := r.Resolver.DB

	if filter != nil {
		if filter.ID != nil {
			query = query.Where("id = ?", *filter.ID)
		}
		if filter.Locale != nil && *filter.Locale != "" {
			query = query.Where("locale = ?", *filter.Locale)
		}
		if filter.ProviderID != nil {
			query = query.Where("provider_id = ?", *filter.ProviderID)
		}
		if filter.FullName != nil && *filter.FullName != "" {
			query = query.Where("full_name ILIKE ?", fmt.Sprintf("%%%s%%", *filter.FullName))
		}
		if filter.IsPremium != nil && !*filter.IsPremium {
			query = query.Where("is_premium = FALSE")
		}
	}
	if *limit > 0 {
		query = query.Limit(*limit)
	}
	if *offset > 0 {
		query = query.Offset(*offset)
	}

	err := query.Order("id").Preload("MetaProvider").Find(&products).Error
	if err != nil {
		log.Println(err)
		return products, err
	}
	return products, nil
}

func (r *queryResolver) ProductsAdminResolver(filter *model.ProductFilter, limit *int, offset *int) ([]*model.ProductAdmin, error) {
	var products []*model.ProductAdmin
	query := r.Resolver.DB

	if filter != nil {
		if filter.ID != nil {
			query = query.Where("id = ?", *filter.ID)
		}
		if filter.Locale != nil && *filter.Locale != "" {
			query = query.Where("locale = ?", *filter.Locale)
		}
		if filter.ProviderID != nil {
			query = query.Where("provider_id = ?", *filter.ProviderID)
		}
		if filter.FullName != nil && *filter.FullName != "" {
			query = query.Where("full_name ILIKE ?", fmt.Sprintf("%%%s%%", *filter.FullName))
		}
		if filter.IsPremium != nil && !*filter.IsPremium {
			query = query.Where("is_premium = FALSE")
		}
	}
	if *limit > 0 {
		query = query.Limit(*limit)
	}
	if *offset > 0 {
		query = query.Offset(*offset)
	}

	err := query.Order("id").Preload("MetaProvider").Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *queryResolver) ProductsByPhoneNumberResolver(phoneNumber model.PhoneNumber, limit *int, offset *int) ([]*model.Product, error) {
	var products []*model.Product
	query := r.Resolver.DB

	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return products, err
	}
	processedCountryCode := reg.ReplaceAllString(phoneNumber.CountryCode, "")
	processedPhoneNumber := reg.ReplaceAllString(phoneNumber.PhoneNumber, "")
	fullPhoneNumber := "+" + processedCountryCode + processedPhoneNumber
	countryISO, err := phonegeocode.New().Country(fullPhoneNumber)
	if err != nil {
		return products, err
	}

	query = query.Where("locale = ?", countryISO)
	if *limit > 0 {
		query = query.Limit(*limit)
	}
	if *offset > 0 {
		query = query.Offset(*offset)
	}

	err = query.Order("id").Preload("MetaProvider").Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *queryResolver) CategoriesResolver(limit *int, offset *int) ([]*model.Category, error) {
	var categories []*model.Category
	query := r.Resolver.DB

	if *limit > 0 {
		query = query.Limit(*limit)
	}
	if *offset > 0 {
		query = query.Offset(*offset)
	}

	err := query.Order("id").Find(&categories).Error
	if err != nil {
		return categories, err
	}
	return categories, nil
}

func (r *queryResolver) ProvidersResolver(filter *model.ProviderFilter, limit *int, offset *int) ([]*model.Provider, error) {
	var providers []*model.Provider
	query := r.Resolver.DB

	if filter != nil && filter.Name != "" {
		query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", filter.Name))
	}
	if *limit > 0 {
		query = query.Limit(*limit)
	}
	if *offset > 0 {
		query = query.Offset(*offset)
	}

	err := query.Order("id").Preload("MetaProvider").Find(&providers).Error
	if err != nil {
		return providers, err
	}
	return providers, nil
}

func (r *queryResolver) CountriesResolver(limit *int, offset *int) ([]*model.Country, error) {
	var countries []*model.Country
	query := r.Resolver.DB

	if *limit > 0 {
		query = query.Limit(*limit)
	}
	if *offset > 0 {
		query = query.Offset(*offset)
	}

	err := query.Order("id").Find(&countries).Error
	if err != nil {
		return countries, err
	}
	return countries, nil
}
