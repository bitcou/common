//go:generate go run github.com/99designs/gqlgen
package graph

import (
	"fmt"
	"github.com/bitcou/bitcou-api/graph/model"
	"github.com/davegardnerisme/phonegeocode"
	"gorm.io/gorm"
	"log"
	"regexp"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *gorm.DB
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
				panic("Check the start and end values for the date range.")
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
				log.Fatal(err)
			}

			if filter.PriceRange.MinPrice != nil && *filter.PriceRange.MinPrice > 0 {
				minPrice = *filter.PriceRange.MinPrice
			}
			if filter.PriceRange.MaxPrice != nil && *filter.PriceRange.MaxPrice > 0 {
				maxPrice = *filter.PriceRange.MaxPrice
			}
			if minPrice > maxPrice || maxPrice < minPrice {
				panic("Check the minimum and maximum values for the price range.")
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

	err := query.Order("id").Find(&purchases).Error
	if err != nil {
		log.Fatal(err)
	}
	return purchases, nil
}

func (r *queryResolver) ClientsResolver(filter *model.ClientFilter, limit *int, offset *int) ([]*model.Client, error) {
	var clients []*model.Client
	query := r.Resolver.DB

	if filter != nil {
		if filter.ID != nil {
			query = query.Where("id = ?", *filter.ID)
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
		log.Fatal(err)
	}
	return clients, nil
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
			query = query.Where("full_name LIKE ?", fmt.Sprintf("%%%s%%", *filter.FullName))
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
		log.Fatal(err)
	}
	return products, nil
}

func (r *queryResolver) ProductsByPhoneNumberResolver(phoneNumber model.PhoneNumber, limit *int, offset *int) ([]*model.Product, error) {
	var products []*model.Product
	query := r.Resolver.DB

	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedCountryCode := reg.ReplaceAllString(phoneNumber.CountryCode, "")
	processedPhoneNumber := reg.ReplaceAllString(phoneNumber.PhoneNumber, "")
	fullPhoneNumber := "+" + processedCountryCode + processedPhoneNumber
	countryISO, err := phonegeocode.New().Country(fullPhoneNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The dialing code of %s is: %v\n", phoneNumber.CountryCode, countryISO)
	fmt.Printf("The full phone number is: %s\n", fullPhoneNumber)
	query = query.Where("locale = ?", countryISO)
	if *limit > 0 {
		query = query.Limit(*limit)
	}
	if *offset > 0 {
		query = query.Offset(*offset)
	}

	err = query.Order("id").Preload("MetaProvider").Find(&products).Error
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
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
		log.Fatal(err)
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
		log.Fatal(err)
	}
	return countries, nil
}
