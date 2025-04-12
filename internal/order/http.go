package main

import (
	"github.com/fhyxzmkh/gorder/order/app"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	app app.Application
}

func (s HTTPServer) PostCustomerCustomerIDOrders(c *gin.Context, customerID string) {
	//todo
}

func (s HTTPServer) GetCustomerCustomerIDOrdersOrderID(c *gin.Context, customerID string, orderID string) {
	//todo
}
