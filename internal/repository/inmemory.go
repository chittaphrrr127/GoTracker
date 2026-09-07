package repository

import (
	"sync"
	"github.com/chittaphrrr127/GoTracker/internal/order"
)

type InMemoryOrderRepo struct {
	orders map[int]order.Order
	mu     sync.Mutex
}

func (r *InMemoryOrderRepo) Add(order order.Order) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[order.ID] = order
}

func (r *InMemoryOrderRepo) GetByID(id int) (order.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	order, ok := r.orders[id]
	if !ok {
		return order.Order{}, order.ErrOrderNotFound //valid zero-value Order. Not nil, because Order isn't a pointer/interface/map/etc.
	}
	return order, nil
}

func (r *InMemoryOrderRepo) Update(order order.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.orders[order.ID]; !ok {
		return order.ErrOrderNotFound
	}
	r.orders[order.ID] = order
	return nil
}

func (r *InMemoryOrderRepo) GetAll() []order.Order {
	r.mu.Lock()
	defer r.mu.Unlock()
	orders := []order.Order{}
	for _, order := range r.orders {
		orders = append(orders, order)
	}
	return orders
}
