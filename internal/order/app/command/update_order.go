package command

import (
	"context"
	"github.com/fhyxzmkh/gorder/common/decorator"
	domain "github.com/fhyxzmkh/gorder/order/domain/order"
	"github.com/sirupsen/logrus"
)

type UpdateOrder struct {
	Order    *domain.Order
	UpdateFn func(context.Context, *domain.Order) (*domain.Order, error)
}

type UpdateOrderHandler decorator.CommandHandler[UpdateOrder, interface{}]

type updateOrderHandler struct {
	orderRepo domain.Repository
}

func NewUpdateOrderHandler(orderRepo domain.Repository, logger *logrus.Entry, client decorator.MetricsClient) UpdateOrderHandler {
	if orderRepo == nil {
		panic("orderRepo is nil")
	}

	return decorator.ApplyCommandDecorators[UpdateOrder, interface{}](
		updateOrderHandler{orderRepo: orderRepo},
		logger,
		client,
	)
}

func (c updateOrderHandler) Handle(ctx context.Context, command UpdateOrder) (interface{}, error) {
	if command.UpdateFn == nil {
		logrus.Warnf("UpdateOrder command is missing UpdateFn function")
		command.UpdateFn = func(_ context.Context, order *domain.Order) (*domain.Order, error) {
			return order, nil
		}
	}

	if err := c.orderRepo.Update(ctx, command.Order, command.UpdateFn); err != nil {
		return nil, err
	}

	return nil, nil
}
