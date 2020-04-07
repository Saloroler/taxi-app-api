package manager

import (
	"taxiapp/cmd/orderapi/application/cache"
	"taxiapp/cmd/orderapi/models"
)

type manager struct {
	ordersHistoryCache cache.OrdersHistoryCache
	ordersCache        cache.OrdersCache
}

type OrderManager interface {
	GetRequestHistory() models.OrdersHistory
	GetRandomOrder() models.OrderTicket
	AddOrderRequestInHistory(orderTicket models.OrderTicket)
}

func NewOrderManager(ordersCache cache.OrdersCache, ordersHistoryCache cache.OrdersHistoryCache) OrderManager {
	return &manager{
		ordersHistoryCache: ordersHistoryCache,
		ordersCache:        ordersCache,
	}
}

func (m *manager) GetRequestHistory() models.OrdersHistory {
	return m.ordersHistoryCache.GetHistory()
}

func (m *manager) AddOrderRequestInHistory(orderTicket models.OrderTicket) {
	m.ordersHistoryCache.SetOrderRequest(orderTicket)
}

func (m *manager) GetRandomOrder() models.OrderTicket {
	return m.ordersCache.GetRandom()
}
