package service

import (
	"context"
	"github.com/fhyxzmkh/gorder/common/metrics"
	"github.com/fhyxzmkh/gorder/order/adapters"
	"github.com/fhyxzmkh/gorder/order/app"
	"github.com/fhyxzmkh/gorder/order/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	orderRepo := adapters.NewMemoryOrderRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	client := metrics.TodoMetrics{}

	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			GetCustomerOrder: query.NewGetCustomerOrderHandler(orderRepo, logger, client),
		},
	}
}
