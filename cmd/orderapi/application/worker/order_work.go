package worker

import (
	"taxiapp/cmd/orderapi/application"
	"taxiapp/cmd/orderapi/application/cache"
)

type UpdateOrders struct {
	ordersCache   cache.OrdersCache
	countOfOrders int
	countOfLoops  int
}

func NewUpdateOrderListJob(ordersCache cache.OrdersCache, countOfOrders int) *UpdateOrders {
	return &UpdateOrders{
		ordersCache:   ordersCache,
		countOfOrders: countOfOrders,
	}
}

func (u *UpdateOrders) Run() error {
	// Removing order from list
	randomOrderIndex := application.GetRandomNumberBetween(1, u.countOfOrders)
	u.ordersCache.RemoveTicketFromOrderList(randomOrderIndex - 1)

	// Adding new ticket in list with checking on uniqueness
	u.ordersCache.AddNewRandomOrderTicket()

	// u.countOfLoops++
	// fmt.Println(fmt.Sprintf("Loop %v, orders: %v", u.countOfLoops, u.ordersCache.GetAll()))
	return nil
}
