package application

import (
	"sync"
	"taxiapp/cmd/orderapi/models"
)

type UpdateOrders struct {
	ordersRWmutex *sync.RWMutex
	orders        *models.OrderList
}

func NewUpdateOrderListJob(orders *models.OrderList, ordersRWMutex *sync.RWMutex) *UpdateOrders {
	return &UpdateOrders{
		ordersRWmutex: ordersRWMutex,
		orders:        orders,
	}
}

func (u *UpdateOrders) Run() error {
	u.ordersRWmutex.Lock()
	defer u.ordersRWmutex.Unlock()

	// Removing order from list
	randomOrderIndex := getRandomNumberBetween(1, 20)
	removeTicketFromList(u.orders, randomOrderIndex)

	// Adding new ticket in list with checking on uniqueness
	for {
		newOrderTicket := generateOrderTicket()
		if !doesOrderTicketAlreadyExist(*u.orders, newOrderTicket) {
			*u.orders = append(*u.orders, newOrderTicket)
			break
		}
	}

	return nil
}

func removeTicketFromList(slices *models.OrderList, orderIndex int) {
	*slices = append((*slices)[:orderIndex], (*slices)[orderIndex+1:]...)
}
