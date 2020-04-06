package api

import (
	"encoding/json"
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

	body, _ := json.Marshal(map[string]string{"key": "test"})
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (api *orderAPI) GetOrdersReport(w http.ResponseWriter, r *http.Request) {

	body, _ := json.Marshal(map[string]string{"key": "test2"})
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
