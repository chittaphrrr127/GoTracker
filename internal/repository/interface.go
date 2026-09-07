package repository

import "github.com/chittaphrrr127/GoTracker/internal/order"

type Repository interface {
	Add(order.Order)
	GetByID(id int) (order.Order, error)
	Update(order order.Order) error
	GetAll() []order.Order
}

