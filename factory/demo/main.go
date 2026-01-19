package main

import (
	"fmt"

	"go-design-patterns/factory"
)

func main() {
	fmt.Println("=== Factory Pattern Demo ===\n")

	// Demonstrate creating different payment processors using the factory
	fmt.Println("1. Creating payment processors via factory:")

	// Create a credit card processor
	ccProcessor, err := factory.CreatePaymentProcessor(
		factory.CreditCard,
		map[string]string{
			"cardNumber": "1234567890123456",
			"cvv":        "123",
		},
	)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("   Created: %s processor\n", ccProcessor.GetName())

	// Create a PayPal processor
	paypalProcessor, err := factory.CreatePaymentProcessor(
		factory.PayPal,
		map[string]string{
			"email": "user@example.com",
		},
	)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("   Created: %s processor\n", paypalProcessor.GetName())

	// Create a bank transfer processor
	bankProcessor, err := factory.CreatePaymentProcessor(
		factory.BankTransfer,
		map[string]string{
			"accountNumber": "987654321",
			"routingNumber": "123456789",
		},
	)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("   Created: %s processor\n\n", bankProcessor.GetName())

	// Demonstrate processing payments
	fmt.Println("2. Processing payments:")
	processors := []factory.PaymentProcessor{ccProcessor, paypalProcessor, bankProcessor}
	amounts := []float64{99.99, 149.50, 299.00}

	for i, processor := range processors {
		fmt.Printf("   Payment %d:\n", i+1)
		if err := processor.Process(amounts[i]); err != nil {
			fmt.Printf("   Error processing payment: %v\n", err)
		}
		fmt.Println()
	}

	// Demonstrate the benefit: easy to add new types without changing existing code
	fmt.Println("3. Factory pattern benefits:")
	fmt.Println("   ✓ Creation logic is centralized in one place")
	fmt.Println("   ✓ Easy to add new payment types (just update the factory)")
	fmt.Println("   ✓ Client code doesn't need to know about concrete types")
	fmt.Println("   ✓ All processors implement the same interface")

	// Show error handling for unknown types
	fmt.Println("\n4. Error handling:")
	_, err = factory.CreatePaymentProcessor("unknown", nil)
	if err != nil {
		fmt.Printf("   ✓ Factory properly handles unknown types: %v\n", err)
	}
}
