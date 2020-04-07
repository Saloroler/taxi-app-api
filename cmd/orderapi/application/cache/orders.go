package cache

import (
	"sync"
	"taxiapp/cmd/orderapi/application"
	"taxiapp/cmd/orderapi/models"
)

type ordersCache struct {
	ordersMutex   *sync.Mutex
	orders        models.OrderList
	countOfOrders int
}

type OrdersCache interface {
	GetAll() models.OrderList
	GetRandom() models.OrderTicket
	UpdateOrders()
}

func NewOrdersCache(countOfOrders int) OrdersCache {
	return &ordersCache{
		ordersMutex:   &sync.Mutex{}, // Initialize mutex to prevent race conditions
		countOfOrders: countOfOrders,
		orders:        application.GenerateUniqueRandomOrders(countOfOrders),
	}
}

// For test purposes
func (c *ordersCache) GetAll() models.OrderList {
	c.ordersMutex.Lock()
	defer c.ordersMutex.Unlock()

	return c.orders
}

func (c *ordersCache) GetRandom() models.OrderTicket {
	c.ordersMutex.Lock()
	defer c.ordersMutex.Unlock()

	randomIndex := application.GetRandomNumberBetween(1, c.countOfOrders)
	return c.orders[randomIndex-1]
}

func (c *ordersCache) UpdateOrders() {
	c.ordersMutex.Lock()
	defer c.ordersMutex.Unlock()

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
