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
	GetAll() models.OrderList
	GetRandom() models.OrderTicket
	UpdateOrders()
}

func NewOrdersCache(ordersRWMutex *sync.RWMutex, countOfOrders int) OrdersCache {
	return &ordersCache{
		ordersRWMutex: ordersRWMutex,
		countOfOrders: countOfOrders,
		orders:        application.GenerateUniqueRandomOrders(countOfOrders),
	}
}

func (c *ordersCache) GetAll() models.OrderList {
	c.ordersRWMutex.RLock()
	defer c.ordersRWMutex.RUnlock()

	return c.orders
}

func (c *ordersCache) GetRandom() models.OrderTicket {
	c.ordersRWMutex.RLock()
	defer c.ordersRWMutex.RUnlock()

	randomIndex := application.GetRandomNumberBetween(1, c.countOfOrders)
	return c.orders[randomIndex-1]
}

func (c *ordersCache) UpdateOrders() {
	c.ordersRWMutex.RLock()
	defer c.ordersRWMutex.RUnlock()

	randomOrderIndex := application.GetRandomNumberBetween(1, c.countOfOrders)
	c.removeTicketFromOrderList(randomOrderIndex - 1)

	// Adding new ticket in list with checking on uniqueness
	c.addNewRandomOrderTicket()
}

func (c *ordersCache) removeTicketFromOrderList(orderIndex int) {
	// Delete with order
	//c.orders = append(c.orders[:orderIndex], c.orders[orderIndex+1:]...)

	// Optimization delete - without preserving order
	c.orders[orderIndex] = c.orders[len(c.orders)-1]
	c.orders = c.orders[:len(c.orders)-1]
}

func (c *ordersCache) addNewRandomOrderTicket() {
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
