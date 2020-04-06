package service

import (
	"bytes"
	"math/rand"
	"taxiapp/cmd/orderapi/models"
	"time"
)

func GenerateUniqueRandomOrders() models.OrderList {
	var orders []models.OrderTicket

	//Will loop as long as in orders will be 50 unique tickets
	for i := 0; i < 50; {
		newOrderTicket := generateOrderTicket()
		doesAlreadyExist := doesOrderTicketAlreadyExist(orders, newOrderTicket)
		if doesAlreadyExist {
			continue
		}
		orders = append(orders, generateOrderTicket())
		i++
	}

	return orders
}

func getRandomNumberBetween(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func generateOrderTicket() models.OrderTicket {
	var b bytes.Buffer
	firtsRandomLetterNumber := getRandomNumberBetween(1, 26)
	b.WriteString(toChar(firtsRandomLetterNumber))

	secondRandomLetterNumber := getRandomNumberBetween(1, 26)
	b.WriteString(toChar(secondRandomLetterNumber))

	return models.OrderTicket(b.String())
}

func toChar(i int) string {
	return string('A' - 1 + 32 + i) // alpabet in lower case 1 - 26
}

func doesOrderTicketAlreadyExist(orders []models.OrderTicket, ticket models.OrderTicket) bool {
	for _, order := range orders {
		if order == ticket {
			return true
		}
	}
	return false
}