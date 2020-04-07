package cache

import (
	"sync"
	"taxiapp/cmd/orderapi/application"
	"taxiapp/cmd/orderapi/models"
)

type ordersCache struct {
	ordersRWMutex *sync.RWMutex
	orders        models.OrderList
	countOfOrders int
}

type OrdersCache interface {
	RemoveTicketFromOrderList(randomOrderIndex int)
	AddNewRandomOrderTicket()
	GetAll() models.OrderList
	GetRandom() models.OrderTicket
}

func NewOrdersCache(ordersRWMutex *sync.RWMutex, countOfOrders int) OrdersCache {
	return &ordersCache{
		ordersRWMutex: ordersRWMutex,
		countOfOrders: countOfOrders,
		orders:        application.GenerateUniqueRandomOrders(countOfOrders),
	}
}

func (c *ordersCache) RemoveTicketFromOrderList(orderIndex int) {
	// Delete with order
	//c.orders = append(c.orders[:orderIndex], c.orders[orderIndex+1:]...)

	// Optimization delete - without preserving order
	c.orders[orderIndex] = c.orders[len(c.orders)-1]
	c.orders = c.orders[:len(c.orders)-1]
}

func (c *ordersCache) AddNewRandomOrderTicket() {
	for {
		newOrderTicket := application.GenerateOrderTicket()
		doesTicketAlreadyExist := c.orders.DoesOrderTicketAlreadyExist(newOrderTicket)
		if doesTicketAlreadyExist {
			continue
		}
		c.orders = append(c.orders, newOrderTicket)
		break
	}
}

func (c *ordersCache) GetAll() models.OrderList {
	return c.orders
}

func (c *ordersCache) GetRandom() models.OrderTicket {
	randomIndex := application.GetRandomNumberBetween(1, c.countOfOrders)
	return c.orders[randomIndex-1]
}
