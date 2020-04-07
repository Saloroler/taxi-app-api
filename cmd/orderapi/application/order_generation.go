package application

import (
	"bytes"
	"math/rand"
	"taxiapp/cmd/orderapi/models"
	"time"
)

func GenerateUniqueRandomOrders(countOfOrders int) models.OrderList {
	var orders models.OrderList

	//Will loop as long as in orders will be 50 unique tickets
	for i := 0; i < countOfOrders; {
		newOrderTicket := GenerateOrderTicket()
		if orders.DoesOrderTicketAlreadyExist(newOrderTicket) {
			continue
		}
		orders = append(orders, GenerateOrderTicket())
		i++
	}

	return orders
}

func GetRandomNumberBetween(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func GenerateOrderTicket() models.OrderTicket {
	var b bytes.Buffer
	firtsRandomLetterNumber := GetRandomNumberBetween(1, 26)
	b.WriteString(toChar(firtsRandomLetterNumber))

	secondRandomLetterNumber := GetRandomNumberBetween(1, 26)
	b.WriteString(toChar(secondRandomLetterNumber))

	return models.OrderTicket(b.String())
}

func toChar(i int) string {
	return string('A' - 1 + 32 + i) // alpabet in lower case 1 - 26
}
