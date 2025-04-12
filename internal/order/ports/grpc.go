package ports

import (
	"context"
	"github.com/fhyxzmkh/gorder/common/genproto/orderpb"
	"github.com/fhyxzmkh/gorder/order/app"
	"github.com/golang/protobuf/ptypes/empty"
)

type GRPCServer struct {
	app app.Application
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func (G GRPCServer) CreateOrder(ctx context.Context, request *orderpb.CreateOrderRequest) (*empty.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCServer) GetOrders(ctx context.Context, request *orderpb.GetOrdersRequest) (*orderpb.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCServer) UpdateOrder(ctx context.Context, order *orderpb.Order) (*empty.Empty, error) {
	//TODO implement me
	panic("implement me")
}
