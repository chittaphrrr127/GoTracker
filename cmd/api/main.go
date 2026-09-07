package main

import (
	"github.com/chittaphrrr127/GoTracker/internal/order"
	"github.com/chittaphrrr127/GoTracker/internal/repository"
	"github.com/chittaphrrr127/GoTracker/internal/service"
)

func main() {
	repo := repository.InMemoryOrderRepo{
		orders: make(map[int]order.Order),
	}

	svc := service.OrderService{repo: &repo}

	repo.Add(order.Order{ID: 1, Customer: "Andrey", Address: "kzn"})
	repo.Add(order.Order{ID: 2, Customer: "Ivan", Address: "msc"})

	svc.DeliverMany(&repo, []int{1, 2})
	svc.PrintAllOrders()
}
