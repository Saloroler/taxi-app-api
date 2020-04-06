package main

import (
	"log"
	"sync"
	"taxiapp/cmd/orderapi/application"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/rakyll/ticktock"
	"github.com/rakyll/ticktock/t"
)

type config struct {
	OrderJobName string `env:"ORDER_JOB_NAME,required"`
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
	fiftyRandomOrders := application.GenerateUniqueRandomOrders()
	orderListJob := application.NewUpdateOrderListJob(&fiftyRandomOrders, ordersMutex)

	// Run every 200 millisecond
	err = ticktock.Schedule(cfg.OrderJobName, orderListJob, &t.When{Every: t.Every(1).Seconds()})
	if err != nil {
		log.Fatal("Failed to schedule cron job for orders", err)
	}
	go ticktock.Start()

}
