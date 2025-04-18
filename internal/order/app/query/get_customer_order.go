package query

import (
	"context"
	"github.com/fhyxzmkh/gorder/common/decorator"
	domain "github.com/fhyxzmkh/gorder/order/domain/order"
	"github.com/sirupsen/logrus"
)

type GetCustomerOrder struct {
	CustomerID string
	OrderID    string
}

type GetCustomerOrderHandler decorator.QueryHandler[GetCustomerOrder, *domain.Order]

type getCustomerOrderHandler struct {
	orderRepo domain.Repository
}

func NewGetCustomerOrderHandler(orderRepo domain.Repository, logger *logrus.Entry, client decorator.MetricsClient) GetCustomerOrderHandler {

	if orderRepo == nil {
		panic("orderRepo is nil")
	}

	return decorator.ApplyQueryDecorators[GetCustomerOrder, *domain.Order](
		getCustomerOrderHandler{orderRepo: orderRepo},
		logger,
		client,
	)
}

func (g getCustomerOrderHandler) Handle(ctx context.Context, query GetCustomerOrder) (*domain.Order, error) {
	o, err := g.orderRepo.Get(ctx, query.OrderID, query.CustomerID)

	if err != nil {
		return nil, err
	}

	return o, nil
}
