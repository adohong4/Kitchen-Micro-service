package service

import (
	"context"

	"github.com/adohong4/Kitchen-Microservice/services/common/genproto/orders"
)

var orderDb = make([]*orders.Order, 0)

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	orderDb = append(orderDb, order)
	return nil
}
