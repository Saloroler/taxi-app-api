package manager

import (
  "sync"
)


type manager struct{
ordersRWMutex *sync.RWMutex
}

type OrderManager interface {

}


func NewOrderManager(ordersRWMutex *sync.RWMutex) OrderManager{
  return &manager{
    ordersRWMutex: ordersRWMutex,
  }
}
