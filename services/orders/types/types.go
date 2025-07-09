package types

import (
	"context"

	"github.com/adohong4/Kitchen-Microservice/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context) []*orders.Order
}
