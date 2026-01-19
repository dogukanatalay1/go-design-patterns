# Factory Pattern

## What's the Point?

You know that feeling when you're writing code and you need to create an object, but you don't know exactly which type until runtime? Or maybe you have a bunch of similar objects that need to be created, but the creation logic is getting messy and scattered all over your code?

That's where the Factory pattern comes in. It's like having a smart helper that knows how to create the right object for you, so you don't have to worry about the details.

## Real-World Scenarios

Think about these situations:

- **Payment processing**: You accept credit cards, PayPal, and bank transfers. Instead of having `if paymentType == "credit"` scattered everywhere, you have a factory that creates the right payment processor.

- **Database connections**: You support MySQL, PostgreSQL, and SQLite. The factory creates the right database connection based on what you need.

- **UI components**: You're building a cross-platform app. The factory creates Windows buttons, Mac buttons, or Linux buttons depending on the OS.

- **Logging**: You want to log to files, console, or a remote service. The factory gives you the right logger.

Basically, if you find yourself doing a lot of `if type == X then create Y` logic, a factory can clean that up.

## The Problem Without Factory

Let's say you're building a payment system. Without a factory, you might write:

```go
func ProcessPayment(method string, amount float64) {
    if method == "credit" {
        processor := &CreditCardProcessor{}
        processor.Process(amount)
    } else if method == "paypal" {
        processor := &PayPalProcessor{}
        processor.Process(amount)
    } else if method == "bank" {
        processor := &BankTransferProcessor{}
        processor.Process(amount)
    }
    // ... and this gets called from 10 different places
}
```

Problems with this:
- The creation logic is duplicated everywhere
- If you add a new payment method, you have to update every place
- Hard to test (can't easily swap implementations)
- Violates the Open/Closed Principle (open for extension, closed for modification)

## How Factory Solves It

The Factory pattern gives you a single place to handle object creation. You ask the factory "give me a payment processor for credit cards" and it handles all the details. Your code becomes:

```go
processor := PaymentFactory.Create("credit")
processor.Process(amount)
```

Much cleaner! And if you need to add a new type, you just update the factory - the rest of your code doesn't change.

## When Should You Use It?

**Good times to use Factory:**
- You don't know the exact type of object until runtime
- You have multiple similar objects that share an interface
- Object creation is complex (lots of setup, validation, etc.)
- You want to centralize creation logic
- You need to swap implementations easily (great for testing)

**Maybe think twice if:**
- Object creation is super simple (just `&MyStruct{}`)
- You only have one type (no need for a factory)
- The creation logic is already simple and not duplicated

## The Basic Idea

1. **Product Interface**: Define what all your objects have in common (they all implement the same interface)

2. **Concrete Products**: The actual types you want to create (CreditCard, PayPal, etc.)

3. **Factory**: A function or struct that knows how to create the right product based on some input

In Go, factories are often just functions that return an interface. Simple and effective!

## Bottom Line

Factory pattern is about delegating object creation to a dedicated place. Instead of your code knowing how to create every possible type, it just asks the factory "give me one of these" and the factory figures out the details. It makes your code cleaner, easier to extend, and easier to test.

Let's build one together step by step!
