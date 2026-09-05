package main

import (
	"errors"
	"fmt"
)

type Order struct {
	ID          int
	Customer    string
	Address     string
	IsDelivered bool
}

func (o *Order) MarkDelivered() {
	o.IsDelivered = true
}

var ErrOrderNotFound = errors.New("order not found")

type OrderRepository interface {
	Add(order Order)
	GetByID(id int) (Order, error)
	Update(order Order) error
	GetAll() []Order
}

type InMemoryOrderRepo struct {
	orders map[int]Order
}

func (i *InMemoryOrderRepo) Add(order Order) {
	i.orders[order.ID] = order
}

func (i *InMemoryOrderRepo) GetByID(id int) (Order, error) {
	order, ok := i.orders[id]
	if !ok {
		return Order{}, ErrOrderNotFound //valid zero-value Order. Not nil, because Order isn't a pointer/interface/map/etc.
	}
	return order, nil
}

func (i *InMemoryOrderRepo) Update(order Order) error { //
	if _, ok := i.orders[order.ID]; !ok {
		return ErrOrderNotFound
	}
	i.orders[order.ID] = order
	return nil
}

func (i *InMemoryOrderRepo) GetAll() []Order {
	orders := []Order{}
	for _, order := range i.orders {
		orders = append(orders, order)
	}
	return orders
}

func PrintAllOrders(repo OrderRepository) {
	orders := repo.GetAll()
	for _, order := range orders {
		fmt.Printf("ID: %d, %s, Delivered: %t\n", order.ID, order.Customer, order.IsDelivered)
	}
}

func main() { //
	repo := InMemoryOrderRepo{
		orders: make(map[int]Order),
	}

	order1 := Order{
		ID:          1,
		Customer:    "Andrey",
		Address:     "a",
		IsDelivered: false,
	}

	order2 := Order{
		ID:          2,
		Customer:    "Ivan",
		Address:     "i",
		IsDelivered: false,
	}

	repo.Add(order1)
	repo.Add(order2)

	fmt.Println("[Before]")
	PrintAllOrders(&repo)

	id := 1

	order, err := repo.GetByID(id)
	if err != nil {
		fmt.Println("Ошибка", err)
	} else {
		order.MarkDelivered()
		if err := repo.Update(order); err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Printf("\nUpdated order %d\n", order.ID)
		}
	}

	fmt.Println("\n[After]")
	PrintAllOrders(&repo)

	order3 := Order{
		ID:          3,
		Customer:    "Assel",
		Address:     "q",
		IsDelivered: false,
	}
	erru := repo.Update(order3)
	if erru != nil {
		fmt.Println("\nОшибка:", erru)
	}
}
