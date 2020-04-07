package api

import (
	"encoding/json"
	"log"
	"net/http"
	"taxiapp/cmd/orderapi/application/manager"
)

type OrderAPI interface {
	GetOrder(w http.ResponseWriter, r *http.Request)
	GetOrdersReport(w http.ResponseWriter, r *http.Request)
}

type orderAPI struct {
	orderManager manager.OrderManager
}

func NewOrderApi(orderManager manager.OrderManager) OrderAPI {
	return &orderAPI{
		orderManager: orderManager,
	}
}

func (api *orderAPI) GetOrder(w http.ResponseWriter, r *http.Request) {
	randomOrder := api.orderManager.GetRandomOrder()
	body, err := json.Marshal(map[string]string{"order": string(randomOrder)})
	if err != nil {
		log.Println("Failed to marshal order, err:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
	//Put request in history
	api.orderManager.AddOrderRequestInHistory(randomOrder)
}

func (api *orderAPI) GetOrdersReport(w http.ResponseWriter, r *http.Request) {
	requests := api.orderManager.GetRequestHistory()
	if len(requests) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	body, err := json.Marshal(requests)
	if err != nil {
		log.Println("Failed to marshal requests history, err:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
