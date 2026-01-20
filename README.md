# Go Design Patterns

A comprehensive collection of design patterns implemented in Go, with detailed explanations and practical examples.

## Learning Approach

This repository teaches design patterns through a structured approach:

1. **Use-case**: Understanding when and why to use each pattern
2. **Implementation**: Step-by-step guide on how to implement the pattern in Go
3. **Examples**: Complete, runnable code examples demonstrating the pattern in action

Each pattern is organized in its own folder with:
- `README.md` - Detailed explanation of the pattern, use-cases, and implementation approach
- `example.go` - Complete implementation example
- `main.go` - Runnable demonstration (if applicable)

## Table of Contents

### Creational Patterns

- [Singleton Pattern](./singleton/) - Ensure a class has only one instance and provide a global point of access to it
- [Factory Pattern](./factory/) - Create objects without specifying the exact class of object that will be created
- [Builder Pattern](./builder/) - Construct complex objects step by step with a fluent interface

### Structural Patterns

_(Coming soon)_

### Behavioral Patterns

_(Coming soon)_

## How to Use

1. Navigate to the pattern folder you want to learn
2. Read the `README.md` for use-case and implementation details
3. Study the `example.go` to see the implementation
4. Run the demo (if available) to see the pattern in action:
   ```bash
   cd singleton/demo
   go run main.go
   
   # or for factory pattern
   cd factory/demo
   go run main.go
   
   # or for builder pattern
   cd builder/demo
   go run main.go
   ```

## Contributing

This is a learning repository. Each pattern is added one by one with comprehensive explanations.
