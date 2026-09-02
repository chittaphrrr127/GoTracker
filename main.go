package main

import "fmt"

type Order struct { //+
	ID          int
	Customer    string
	Address     string
	IsDelivered bool
}

type InMemoryOrderRepo struct { //+
	orders map[int]Order
}

type OrderRepository interface { //+
	Add(order Order)
	GetByID(id int) (Order, bool)
	Update(order Order)
	GetAll() (order []Order)
}

func (o *Order) MarkDelivered() { //+
	o.IsDelivered = true
}

func (i *InMemoryOrderRepo) Add(order Order) { //+
	i.orders[order.ID] = order
}

func (i *InMemoryOrderRepo) GetByID(id int) (Order, bool) { //+
	order, ok := i.orders[id]
	return order, ok
}

func (i *InMemoryOrderRepo) Update(order Order) { //+
	i.orders[order.ID] = order
}

func (i *InMemoryOrderRepo) GetAll() []Order { //+
	orders := []Order{}
	for _, order := range i.orders {
		orders = append(orders, order)
	}
	return orders
}

func PrintAllOrders(repo OrderRepository) { //+
	orders := repo.GetAll()
	for _, order := range orders {
		fmt.Printf("ID: %d, %s, Delivered: %v\n", order.ID, order.Customer, order.IsDelivered)
	}
}

func NewInMemoryOrderRepo() *InMemoryOrderRepo { //?
	return &InMemoryOrderRepo{
		orders: make(map[int]Order),
	}
}

func main() {
	// repo := InMemoryOrderRepo{
	// 	orders: make(map[int]Order),
	// }

	repo:=NewInMemoryOrderRepo()

	order1 := Order{
		ID:          1,
		Customer:    "Andrey",
		Address:     "A1",
		IsDelivered: false,
	}

	order2 := Order{
		ID:          2,
		Customer:    "Ivan",
		Address:     "B1",
		IsDelivered: false,
	}

	repo.Add(order1)
	repo.Add(order2)

	fmt.Println("[Before]")
	PrintAllOrders(repo)

	order, exists := repo.GetByID(1)
	if exists {
		order.MarkDelivered()
		repo.Update(order)
	}

	fmt.Println("[After]")
	PrintAllOrders(repo)
}
