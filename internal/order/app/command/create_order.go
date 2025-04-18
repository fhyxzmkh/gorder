package command

import (
	"context"
	"github.com/fhyxzmkh/gorder/common/decorator"
	"github.com/fhyxzmkh/gorder/common/genproto/orderpb"
	domain "github.com/fhyxzmkh/gorder/order/domain/order"
	"github.com/sirupsen/logrus"
)

type CreateOrder struct {
	CustomerID string
	Items      []*orderpb.ItemWithQuantity
}

type CreateOrderResult struct {
	OrderID string
}

type CreateOrderHandler decorator.CommandHandler[CreateOrder, *CreateOrderResult]

type createOrderHandler struct {
	orderRepo domain.Repository
}

func NewCreateOrderHandler(orderRepo domain.Repository, logger *logrus.Entry, client decorator.MetricsClient) CreateOrderHandler {
	if orderRepo == nil {
		panic("orderRepo is nil")
	}

	return decorator.ApplyCommandDecorators[CreateOrder, *CreateOrderResult](
		createOrderHandler{orderRepo: orderRepo},
		logger,
		client,
	)
}

func (c createOrderHandler) Handle(ctx context.Context, command CreateOrder) (*CreateOrderResult, error) {
	//TODO call stock grpc to get items

	var stockResponse []*orderpb.Item

	for _, item := range command.Items {
		stockResponse = append(stockResponse, &orderpb.Item{
			ID:       item.ID,
			Quantity: item.Quantity,
		})
	}

	o, err := c.orderRepo.Create(ctx, &domain.Order{
		CustomerID: command.CustomerID,
		Items:      stockResponse,
	})

	if err != nil {
		return nil, err
	}

	return &CreateOrderResult{
		OrderID: o.ID,
	}, nil
}
