package main

import (
	"fmt"
	"sync"
)

// Getters and Setters Advantages
//		Allow further functionality to be added later, validation, thread safety etc.
//		Hide internals
//		Provide a nice debugging point

// Customer struct demonstrates encapsulation and balance handling.
type Customer struct {
	name    string
	balance float64
	mu      sync.Mutex
}

// Balance is a getter method that provides access to the balance.
// Getter method follows the idiomatic Go convention: no "Get" prefix.
func (c *Customer) Balance() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.balance
}

// SetBalance is a setter method that allows updating the balance.
// It includes a behavior to ensure the balance cannot go below zero.
func (c *Customer) SetBalance(amount float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if amount < 0 {
		c.balance = 0
	} else {
		c.balance = amount
	}
}

// NewCustomer is a constructor to create a new customer.
func NewCustomer(name string, balance float64) *Customer {
	return &Customer{name: name, balance: balance}
}

func main() {
	// Creating a new customer
	customer := NewCustomer("John Doe", 100.0)

	// Using the getter to check the balance
	currentBalance := customer.Balance()
	fmt.Printf("Current Balance: %.2f\n", currentBalance)

	// Using the setter to adjust the balance
	if currentBalance < 0 {
		customer.SetBalance(0)
	} else {
		customer.SetBalance(currentBalance - 50)
	}

	// Display the updated balance
	fmt.Printf("Updated Balance: %.2f\n", customer.Balance())
}
