package app

import "github.com/fhyxzmkh/gorder/order/app/query"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}
type Queries struct {
	GetCustomerOrder query.GetCustomerOrderHandler
}
