package worker

import (
	"taxiapp/cmd/orderapi/application/cache"
)

type UpdateOrders struct {
	ordersCache cache.OrdersCache
}

func NewUpdateOrderListWorker(ordersCache cache.OrdersCache) *UpdateOrders {
	return &UpdateOrders{
		ordersCache: ordersCache,
	}
}

// Realized interface in tick library
func (u *UpdateOrders) Run() error {
	u.ordersCache.UpdateOrders()

	return nil
}
