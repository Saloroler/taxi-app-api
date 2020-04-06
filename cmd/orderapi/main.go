package main

import (
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
	"taxiapp/cmd/orderapi/adapters/api"
	"taxiapp/cmd/orderapi/application/manager"
	"taxiapp/cmd/orderapi/application/service"

	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rakyll/ticktock"
	"github.com/rakyll/ticktock/t"
)

type config struct {
	OrderJobName string `env:"ORDER_JOB_NAME,required"`
	Port         int    `env:"PORT,required"`
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

	// Initialize mutex to prevent race conditions
	ordersMutex := &sync.RWMutex{}

	// Initialize orders and job
	fiftyRandomOrders := service.GenerateUniqueRandomOrders()
	orderListJob := service.NewUpdateOrderListJob(&fiftyRandomOrders, ordersMutex)

	// Run every 200 millisecond
	err = ticktock.Schedule(cfg.OrderJobName, orderListJob, &t.When{Every: t.Every(1).Seconds()})
	if err != nil {
		log.Fatal("Failed to schedule cron job for orders", err)
	}
	go ticktock.Start()

	orderManager := manager.NewOrderManager(ordersMutex)

	orderController := api.NewOrderApi(orderManager)

	router := mux.NewRouter()
	adminRouter := router.PathPrefix("/admin").Subrouter()

	router.HandleFunc("/order", orderController.GetOrder).Methods(http.MethodGet)
	adminRouter.HandleFunc("/orders", orderController.GetOrdersReport).Methods(http.MethodGet)

	tcpAddr := net.TCPAddr{Port: cfg.Port}
	log.Println("Server is starting on port", cfg.Port)
	if err := http.ListenAndServe(tcpAddr.String(), router); err != nil {
		log.Fatal("Failed to listen port:"+strconv.Itoa(cfg.Port), err.Error())
	}
}
