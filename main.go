package main

import "fmt"

func main() {
	// var (
	// 	orderID       int
	// 	customerName  string
	// 	isDelivered   bool
	// 	isReadyToShip bool
	// 	orderIDs      []int
	// )

	var orderID int = 1001
	var customerName string = "Ivan"
	var isDelivered bool = false
	fmt.Println("Order ID:", orderID)
	fmt.Println("Customer Name:", customerName)
	fmt.Println("Is Delivered:", isDelivered)

	orderIDs := []int{}
	// for i := 1; i <= 3; i++ {
	// 	orderIDs = append(orderIDs, i+100)
	// }
	orderIDs = append(orderIDs, 101, 102, 103)
	fmt.Println("orderIDs:", orderIDs)
	fmt.Println("len(orderIDs)=", len(orderIDs), "cap(orderIDs)=", cap(orderIDs))

	orderCount := map[string]int{
		"Alice": 5,
		"Bob":   1,
	}
	// orderCount["Alice"] = 5
	// orderCount["Bob"] = 1
	fmt.Println("orderCount:", orderCount)

	// fmt.Println("Is Alice ready to ship?")
	// isReadyToShip = false
	// if len(orderCount) > 2 {
	// 	isReadyToShip = true
	// }
	// fmt.Println(isReadyToShip)

	isReadyToShip:= orderCount["Alice"]>2
	fmt.Println("Is Alice ready to ship?", isReadyToShip)
}
