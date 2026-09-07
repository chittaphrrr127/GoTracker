package base

import (
	"errors"
	"fmt"
	"sync"
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
	mu     sync.Mutex
}

func (r *InMemoryOrderRepo) Add(order Order) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[order.ID] = order
}

func (r *InMemoryOrderRepo) GetByID(id int) (Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	order, ok := r.orders[id]
	if !ok {
		return Order{}, ErrOrderNotFound //valid zero-value Order. Not nil, because Order isn't a pointer/interface/map/etc.
	}
	return order, nil
}

func (r *InMemoryOrderRepo) Update(order Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.orders[order.ID]; !ok {
		return ErrOrderNotFound
	}
	r.orders[order.ID] = order
	return nil
}

func (r *InMemoryOrderRepo) GetAll() []Order {
	r.mu.Lock()
	defer r.mu.Unlock()
	orders := []Order{}
	for _, order := range r.orders {
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

func DeliverMany(repo OrderRepository, ids []int) {
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func(orderID int) {
			defer wg.Done()
			order, err := repo.GetByID(orderID)
			if err != nil {
				fmt.Printf("Ошибка: %v (ID: %d)\n", err, orderID)
				return //!
			}
			order.MarkDelivered()
			if err := repo.Update(order); err != nil { //===============erru
				fmt.Printf("Ошибка при обновлении ID %d: %v\n", orderID, err)
				return //!
			}
			fmt.Printf("Order %d delivered\n", orderID)
		}(id)
	}
	wg.Wait()
}

func main() {
	repo := InMemoryOrderRepo{
		orders: make(map[int]Order),
	}

	repo.Add(Order{ID: 1, Customer: "Andrey", Address: "kzn"})
	repo.Add(Order{ID: 2, Customer: "ivan", Address: "msc"})
	repo.Add(Order{ID: 3, Customer: "Katya", Address: "smr"})

	fmt.Println("[До доставки]")
	PrintAllOrders(&repo)

	fmt.Println("\n[Результаты доставки]")
	DeliverMany(&repo, []int{1, 2, 999})

	fmt.Println("\n[После доставки]")
	PrintAllOrders(&repo)
}
