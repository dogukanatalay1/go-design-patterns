package main

import (
	"fmt"
	"time"

	"go-design-patterns/builder"
)

func main() {
	fmt.Println("=== Builder Pattern Demo ===\n")

	// Demonstrate building a server config with only required fields
	fmt.Println("1. Building a minimal server config:")
	minimalConfig, err := builder.NewServerConfigBuilder().
		Host("localhost").
		Port(8080).
		Build()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("   Host: %s, Port: %d, SSL: %v\n", minimalConfig.Host, minimalConfig.Port, minimalConfig.SSL)
	fmt.Println("   ✓ Only set what we need, defaults applied for the rest\n")

	// Demonstrate building a full-featured config
	fmt.Println("2. Building a full-featured server config:")
	fullConfig, err := builder.NewServerConfigBuilder().
		Host("api.example.com").
		Port(443).
		EnableSSL(true).
		Timeout(60 * time.Second).
		MaxConnections(1000).
		ReadTimeout(30 * time.Second).
		WriteTimeout(30 * time.Second).
		DatabaseURL("postgresql://localhost:5432/mydb").
		EnableCache(true).
		LogLevel("debug").
		Build()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("   Host: %s\n", fullConfig.Host)
	fmt.Printf("   Port: %d\n", fullConfig.Port)
	fmt.Printf("   SSL: %v\n", fullConfig.SSL)
	fmt.Printf("   Timeout: %v\n", fullConfig.Timeout)
	fmt.Printf("   Max Connections: %d\n", fullConfig.MaxConnections)
	fmt.Printf("   Database URL: %s\n", fullConfig.DatabaseURL)
	fmt.Printf("   Cache Enabled: %v\n", fullConfig.CacheEnabled)
	fmt.Printf("   Log Level: %s\n", fullConfig.LogLevel)
	fmt.Println("   ✓ Readable, step-by-step construction\n")

	// Demonstrate partial configuration
	fmt.Println("3. Building a config with some custom settings:")
	partialConfig, err := builder.NewServerConfigBuilder().
		Host("staging.example.com").
		Port(3000).
		EnableSSL(true).
		LogLevel("warn").
		Build()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("   Host: %s, Port: %d, SSL: %v, Log Level: %s\n",
		partialConfig.Host, partialConfig.Port, partialConfig.SSL, partialConfig.LogLevel)
	fmt.Printf("   Timeout (default): %v\n", partialConfig.Timeout)
	fmt.Println("   ✓ Mix of custom and default values\n")

	// Demonstrate validation
	fmt.Println("4. Demonstrating validation:")

	// Missing required field
	_, err = builder.NewServerConfigBuilder().
		Port(8080).
		Build()
	if err != nil {
		fmt.Printf("   ✓ Validation caught missing host: %v\n", err)
	}

	// Invalid port
	_, err = builder.NewServerConfigBuilder().
		Host("localhost").
		Port(99999).
		Build()
	if err != nil {
		fmt.Printf("   ✓ Validation caught invalid port: %v\n", err)
	}

	// Invalid log level
	_, err = builder.NewServerConfigBuilder().
		Host("localhost").
		Port(8080).
		LogLevel("invalid").
		Build()
	if err != nil {
		fmt.Printf("   ✓ Validation caught invalid log level: %v\n", err)
	}

	fmt.Println("\n5. Builder pattern benefits:")
	fmt.Println("   ✓ Readable: Each field is clearly labeled")
	fmt.Println("   ✓ Flexible: Set only what you need")
	fmt.Println("   ✓ Safe: Validation before object creation")
	fmt.Println("   ✓ Fluent: Natural method chaining")
	fmt.Println("   ✓ Defaults: Sensible defaults for optional fields")
}
