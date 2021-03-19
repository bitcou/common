package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/bitcou/common/auth"
	"github.com/bitcou/common/dbmodels/graph/generated"
	"github.com/bitcou/common/dbmodels/graph/model"
	commonErrors "github.com/bitcou/common/errors"
)

func (r *queryResolver) Clients(ctx context.Context, filter *model.ClientFilter, limit *int, offset *int) ([]*model.Client, error) {
	// TODO if isAdmin return all clients
	clientInfo := auth.ForContext(ctx)
	if clientInfo == nil {
		log.Println("no client found")
		return nil, errors.New("no client info")
	}
	// TODO if is admin return all client info.
	log.Println("clientInfo", clientInfo)
	if filter == nil {
		filter = &model.ClientFilter{
			ID: &clientInfo.ID,
		}
	} else {
		filter.ID = &clientInfo.ID
	}
	return r.ClientsResolver(filter, limit, offset)
}

func (r *queryResolver) Categories(ctx context.Context, limit *int, offset *int) ([]*model.Category, error) {
	return r.CategoriesResolver(limit, offset)
}

func (r *queryResolver) Purchases(ctx context.Context, filter *model.PurchaseFilter, limit *int, offset *int) ([]*model.Purchase, error) {
	clientInfo := auth.ForContext(ctx)
	if clientInfo == nil {
		log.Println("no client found")
		return nil, errors.New("no client info")
	}
	// TODO if is admin return all client info.
	log.Println("clientInfo", clientInfo)
	if filter == nil {
		filter = &model.PurchaseFilter{
			ClientID: &clientInfo.ID,
		}
	} else {
		filter.ClientID = &clientInfo.ID
	}
	return r.PurchasesResolver(filter, limit, offset)
}

func (r *queryResolver) Products(ctx context.Context, filter *model.ProductFilter, limit *int, offset *int) ([]*model.Product, error) {
	clientInfo := auth.ForContext(ctx)
	if clientInfo == nil {
		log.Println("no client found")
		return nil, errors.New("no client info")
	}
	log.Println("clientInfo", clientInfo)
	if filter == nil {
		filter = &model.ProductFilter{
			IsPremium: &clientInfo.IsPremium,
		}
	} else {
		filter.IsPremium = &clientInfo.IsPremium
	}

	return r.ProductsResolver(filter, limit, offset)
}

func (r *queryResolver) ProductsAdmin(ctx context.Context, filter *model.ProductFilter, limit *int, offset *int) ([]*model.ProductAdmin, error) {
	clientInfo := auth.ForContext(ctx)
	if clientInfo == nil {
		return nil, errors.New("no client info")
	}
	if !clientInfo.IsPremium {
		return nil, commonErrors.ErrorAdminOnly
	}
	fmt.Println("before procesiing", filter)
	fmt.Println("clientInfo", clientInfo)

	return r.ProductsAdminResolver(filter, limit, offset)
}

func (r *queryResolver) Providers(ctx context.Context, filter *model.ProviderFilter, limit *int, offset *int) ([]*model.Provider, error) {
	return r.ProvidersResolver(filter, limit, offset)
}

func (r *queryResolver) Countries(ctx context.Context, limit *int, offset *int) ([]*model.Country, error) {
	return r.CountriesResolver(limit, offset)
}

func (r *queryResolver) ProductsByPhoneNumber(ctx context.Context, phoneNumber model.PhoneNumber, limit *int, offset *int) ([]*model.Product, error) {
	return r.ProductsByPhoneNumberResolver(phoneNumber, limit, offset)
}

func (r *queryResolver) AccountInfo(ctx context.Context, username string, password string) (*model.Client, error) {
	return r.ClientInfoResolver(username, password)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
