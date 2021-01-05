package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wgTake sync.WaitGroup
var wgCook sync.WaitGroup
var wgBring sync.WaitGroup

// Simple structure which tags chef and order numbe
type chefOrder struct {
	chef  string
	order int
}

// Same as above, duplex channel
type chefCooked struct {
	chef  string
	order int
}

// Function to select a Waiter
func selectWaiter() string {
	allWaiters := []string{"A", "B", "C"}
	randomW := rand.Intn(len(allWaiters))
	wpick := allWaiters[randomW]
	return wpick
}

// Function to select a Chef
func selectChef() string {
	allChefs := []string{"Jack", "Bob", "Mark"}
	randomC := rand.Intn(len(allChefs))
	cpick := allChefs[randomC]
	return cpick
}

// Function where the waiter takes order
func takeOrder(order int, orderChan chan chefOrder, cookChan chan chefCooked) {
	defer wgTake.Done()
	chef := selectChef()
	waiter := selectWaiter()

	fmt.Printf("Waiter %s takes order %d to chef %s \n", waiter, order, chef)
	orderChan <- chefOrder{chef, order}
}

// Function where the waiter brings order
func bringOrder(orderChan chan chefOrder, cookChan chan chefCooked) {
	defer wgBring.Done()
	waiter := selectWaiter()
	cookedOrder := <-cookChan
	fmt.Printf("Waiter %s brings order %d from chef %s \n", waiter, cookedOrder.order, cookedOrder.chef)
}

// Function where the chef cooks order
func cookOrder(orderChan chan chefOrder, cookChan chan chefCooked) {
	defer wgCook.Done()
	gotOrder := <-orderChan
	fmt.Printf("Chef %s cooks order %d \n", gotOrder.chef, gotOrder.order)
	cookChan <- chefCooked{gotOrder.chef, gotOrder.order}
}

// Main function
func main() {
	totalOrders := 5

	orderChan := make(chan chefOrder)
	cookChan := make(chan chefCooked)

	for i := 0; i < totalOrders; i++ {
		wgTake.Add(1)
		wgCook.Add(1)
		wgBring.Add(1)
		go takeOrder(i, orderChan, cookChan)
		go cookOrder(orderChan, cookChan)
		go bringOrder(orderChan, cookChan)
	}

	wgTake.Wait()
	close(orderChan)
	wgCook.Wait()
	close(cookChan)
	wgBring.Wait()
}
