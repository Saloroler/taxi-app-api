package cache

import (
	"sync"
	"taxiapp/cmd/orderapi/models"
)

type ordersHistoryCache struct {
	ordersRWMutex   *sync.RWMutex
	historyOrderMap models.OrdersHistory
}

type OrdersHistoryCache interface {
	GetHistory() models.OrdersHistory
	SetOrderRequest(orderTicket models.OrderTicket)
}

func NewOrderHistoryCache(ordersRWMutex *sync.RWMutex) OrdersHistoryCache {
	return &ordersHistoryCache{
		ordersRWMutex:   ordersRWMutex,
		historyOrderMap: make(models.OrdersHistory),
	}
}

func (h *ordersHistoryCache) GetHistory() models.OrdersHistory {
	return h.historyOrderMap
}

func (h *ordersHistoryCache) SetOrderRequest(orderTicket models.OrderTicket) {
	h.historyOrderMap[orderTicket]++
}
