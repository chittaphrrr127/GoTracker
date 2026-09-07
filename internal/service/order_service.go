package service

import (
	"fmt"
	"sync"
	"github.com/chittaphrrr127/GoTracker/internal/order"
	"github.com/chittaphrrr127/GoTracker/internal/repository"
)

type OrderService struct {
	repo repository.Repository
}

func DeliverMany(o *OrderService, ids []int) {
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func(orderID int) {
			defer wg.Done()
			order, err := o.repo.GetByID(orderID)
			if err != nil {
				fmt.Printf("Ошибка: %v (ID: %d)\n", err, orderID)
				return //!
			}
			order.MarkDelivered()
			if err := o.repo.Update(order); err != nil { //===============erru
				fmt.Printf("Ошибка при обновлении ID %d: %v\n", orderID, err)
				return //!
			}
			fmt.Printf("Order %d delivered\n", orderID)
		}(id)
	}
	wg.Wait()
}

func (o *OrderService) PrintAllOrders() {
	for _, order := range o.repo.GetAll() {
		fmt.Printf("ID: %d, %s, Delivered: %t\n", order.ID, order.Customer, order.IsDelivered)
	}
}
