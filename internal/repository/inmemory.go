package repository

import (
	"sync"

	"github.com/chittaphrrr127/GoTracker/internal/order"
)

type InMemoryOrderRepo struct {
	orders map[int]order.Order
	mu     sync.Mutex
}

func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{
		orders: make(map[int]order.Order),
	}
}

func (r *InMemoryOrderRepo) Add(o order.Order) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[o.ID] = o
}

func (r *InMemoryOrderRepo) GetByID(id int) (order.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	ord, ok := r.orders[id]
	if !ok {
		return order.Order{}, order.ErrOrderNotFound //valid zero-value Order. Not nil, because Order isn't a pointer/interface/map/etc.
	}
	return ord, nil
}

func (r *InMemoryOrderRepo) Update(o order.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.orders[o.ID]; !ok {
		return order.ErrOrderNotFound
	}
	r.orders[o.ID] = o
	return nil
}

func (r *InMemoryOrderRepo) GetAll() []order.Order {
	r.mu.Lock()
	defer r.mu.Unlock()
	orders := []order.Order{}
	for _, o := range r.orders {
		orders = append(orders, o)
	}
	return orders
}
