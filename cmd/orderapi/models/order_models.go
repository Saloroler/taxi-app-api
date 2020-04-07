package models

type OrderTicket string

type OrderList []OrderTicket

type OrdersHistory map[OrderTicket]int

func (list OrderList) DoesOrderTicketAlreadyExist(ticket OrderTicket) bool {
	for _, order := range list {
		if order == ticket {
			return true
		}
	}
	return false
}
