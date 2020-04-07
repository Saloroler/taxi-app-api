package cache

import (
	"sync"
	"taxiapp/cmd/orderapi/models"
)

type ordersHistoryCache struct {
	historyOrderMap models.OrdersHistory
	historyMutex    *sync.Mutex
}

type OrdersHistoryCache interface {
	GetHistory() models.OrdersHistory
	SetOrderRequest(orderTicket models.OrderTicket)
}

func NewOrderHistoryCache() OrdersHistoryCache {
	return &ordersHistoryCache{
		historyOrderMap: make(models.OrdersHistory),
		historyMutex:    &sync.Mutex{},
	}
}

func (h *ordersHistoryCache) GetHistory() models.OrdersHistory {
	h.historyMutex.Lock()
	defer h.historyMutex.Unlock()

	return h.historyOrderMap
}

func (h *ordersHistoryCache) SetOrderRequest(orderTicket models.OrderTicket) {
	h.historyMutex.Lock()
	defer h.historyMutex.Unlock()

	h.historyOrderMap[orderTicket]++
}
