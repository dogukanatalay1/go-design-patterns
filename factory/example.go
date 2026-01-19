package factory

import "fmt"

// Step 1: Define the Product Interface
// This is what all our products will have in common.
// In our example, we'll create different types of payment processors.
// They all need to be able to process payments, so we define that common behavior here.

type PaymentProcessor interface {
	Process(amount float64) error
	GetName() string
}

// Step 2: Create Concrete Product Types
// These are the actual implementations of our PaymentProcessor interface.
// Each one handles payments differently, but they all implement the same interface.

// CreditCardProcessor handles credit card payments
type CreditCardProcessor struct {
	cardNumber string
	cvv        string
}

func (c *CreditCardProcessor) Process(amount float64) error {
	fmt.Printf("Processing $%.2f via Credit Card ending in %s\n", amount, c.cardNumber[len(c.cardNumber)-4:])
	// Simulate processing logic
	return nil
}

func (c *CreditCardProcessor) GetName() string {
	return "Credit Card"
}

// PayPalProcessor handles PayPal payments
type PayPalProcessor struct {
	email string
}

func (p *PayPalProcessor) Process(amount float64) error {
	fmt.Printf("Processing $%.2f via PayPal for %s\n", amount, p.email)
	// Simulate processing logic
	return nil
}

func (p *PayPalProcessor) GetName() string {
	return "PayPal"
}

// BankTransferProcessor handles bank transfer payments
type BankTransferProcessor struct {
	accountNumber string
	routingNumber string
}

func (b *BankTransferProcessor) Process(amount float64) error {
	fmt.Printf("Processing $%.2f via Bank Transfer to account %s\n", amount, b.accountNumber)
	// Simulate processing logic
	return nil
}

func (b *BankTransferProcessor) GetName() string {
	return "Bank Transfer"
}

// Step 3: Create the Factory Function
// This is where the magic happens! The factory takes some input (like a payment type)
// and returns the appropriate PaymentProcessor.
// This centralizes all the creation logic in one place.

type PaymentType string

const (
	CreditCard   PaymentType = "credit"
	PayPal       PaymentType = "paypal"
	BankTransfer PaymentType = "bank"
)

// CreatePaymentProcessor is our factory function.
// It takes a payment type and returns the appropriate processor.
// Notice how all the "if type == X" logic is here, not scattered everywhere!
func CreatePaymentProcessor(paymentType PaymentType, details map[string]string) (PaymentProcessor, error) {
	switch paymentType {
	case CreditCard:
		return &CreditCardProcessor{
			cardNumber: details["cardNumber"],
			cvv:        details["cvv"],
		}, nil

	case PayPal:
		return &PayPalProcessor{
			email: details["email"],
		}, nil

	case BankTransfer:
		return &BankTransferProcessor{
			accountNumber: details["accountNumber"],
			routingNumber: details["routingNumber"],
		}, nil

	default:
		return nil, fmt.Errorf("unknown payment type: %s", paymentType)
	}
}
