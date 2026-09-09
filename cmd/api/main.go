package main

import (
	"github.com/chittaphrrr127/GoTracker/internal/order"
	"github.com/chittaphrrr127/GoTracker/internal/repository"
	"github.com/chittaphrrr127/GoTracker/internal/service"
)

func main() {
	repo := repository.NewInMemoryOrderRepo()
	svc := service.NewOrderService(repo)

	repo.Add(order.Order{ID: 1, Customer: "jk", Address: "sl"})
	repo.Add(order.Order{ID: 2, Customer: "jm", Address: "ps"})

	svc.DeliverMany([]int{1, 2})
	svc.PrintAllOrders()
}
