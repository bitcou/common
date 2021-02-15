package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/bitcou/common/dbmodels/graph/generated"
	"github.com/bitcou/common/dbmodels/graph/model"
)

func (r *queryResolver) Clients(ctx context.Context, filter *model.ClientFilter, limit *int, offset *int) ([]*model.Client, error) {
	return r.ClientsResolver(filter, limit, offset)
}

func (r *queryResolver) Categories(ctx context.Context, limit *int, offset *int) ([]*model.Category, error) {
	return r.CategoriesResolver(limit, offset)
}

func (r *queryResolver) Purchases(ctx context.Context, filter *model.PurchaseFilter, limit *int, offset *int) ([]*model.Purchase, error) {
	return r.PurchasesResolver(filter, limit, offset)
}

func (r *queryResolver) Products(ctx context.Context, filter *model.ProductFilter, limit *int, offset *int) ([]*model.Product, error) {
	return r.ProductsResolver(filter, limit, offset)
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
