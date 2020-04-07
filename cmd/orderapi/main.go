package main

import (
	"log"
	"net"
	"net/http"
	"strconv"
	"taxiapp/cmd/orderapi/adapters/api"
	"taxiapp/cmd/orderapi/application/cache"
	"taxiapp/cmd/orderapi/application/manager"
	"taxiapp/cmd/orderapi/application/worker"

	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rakyll/ticktock"
	"github.com/rakyll/ticktock/t"
)

type config struct {
	OrderJobName  string `env:"ORDER_JOB_NAME,required"`
	Port          int    `env:"PORT,required"`
	CountOfOrders int    `env:"COUNT_OF_ORDERS,required"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load config,e:", err.Error())
	}

	cfg := config{}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatal("Failed to parse config,e:", err.Error())
	}

	// Initialize orders and history cache
	ordersCache := cache.NewOrdersCache(cfg.CountOfOrders)
	ordersHistory := cache.NewOrderHistoryCache()

	// Manager and workers
	orderManager := manager.NewOrderManager(ordersCache, ordersHistory)
	orderListWorker := worker.NewUpdateOrderListWorker(ordersCache)

	orderApiController := api.NewOrderApi(orderManager)

	// Run every 200 millisecond
	err = ticktock.Schedule(cfg.OrderJobName, orderListWorker, &t.When{Every: t.Every(200).Milliseconds()})
	if err != nil {
		log.Fatal("Failed to schedule cron job for orders", err)
	}
	go ticktock.Start()

	router := mux.NewRouter()
	adminRouter := router.PathPrefix("/admin").Subrouter()

	router.HandleFunc("/order", orderApiController.GetOrder).Methods(http.MethodGet)
	adminRouter.HandleFunc("/orders", orderApiController.GetOrdersReport).Methods(http.MethodGet)

	tcpAddr := net.TCPAddr{Port: cfg.Port}
	log.Println("Server is starting on port", cfg.Port)
	if err := http.ListenAndServe(tcpAddr.String(), router); err != nil {
		log.Fatal("Failed to listen port:"+strconv.Itoa(cfg.Port), err.Error())
	}
}
