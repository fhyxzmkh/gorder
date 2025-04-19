package service

import (
	"context"
	"github.com/fhyxzmkh/gorder/common/metrics"
	"github.com/fhyxzmkh/gorder/order/adapters"
	"github.com/fhyxzmkh/gorder/order/app"
	"github.com/fhyxzmkh/gorder/order/app/command"
	"github.com/fhyxzmkh/gorder/order/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	orderRepo := adapters.NewMemoryOrderRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	client := metrics.TodoMetrics{}

	return app.Application{
		Commands: app.Commands{
			CreateOrder: command.NewCreateOrderHandler(orderRepo, logger, client),
			UpdateOrder: command.NewUpdateOrderHandler(orderRepo, logger, client),
		},
		Queries: app.Queries{
			GetCustomerOrder: query.NewGetCustomerOrderHandler(orderRepo, logger, client),
		},
	}
}
