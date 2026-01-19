package main

import (
	"fmt"
	"sync"
	"time"

	"go-design-patterns/singleton"
)

func main() {
	fmt.Println("=== Singleton Pattern Demo ===\n")

	// Demonstrate that multiple calls to GetInstance() return the same instance
	fmt.Println("1. Getting multiple instances:")
	db1 := singleton.GetInstance()
	db2 := singleton.GetInstance()
	db3 := singleton.GetInstance()

	fmt.Printf("   db1 connection ID: %d\n", db1.GetConnectionID())
	fmt.Printf("   db2 connection ID: %d\n", db2.GetConnectionID())
	fmt.Printf("   db3 connection ID: %d\n", db3.GetConnectionID())

	if db1 == db2 && db2 == db3 {
		fmt.Println("   ✓ All references point to the same instance!\n")
	}

	// Demonstrate usage
	fmt.Println("2. Using the singleton instance:")
	db1.Connect()
	db1.Query("SELECT * FROM users")
	db1.Query("SELECT * FROM products")
	db1.Disconnect()

	fmt.Println()

	// Demonstrate thread-safety with concurrent access
	fmt.Println("3. Testing thread-safety with concurrent access:")
	var wg sync.WaitGroup
	instances := make([]*singleton.DatabaseConnection, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			// Small delay to increase chance of concurrent access
			time.Sleep(time.Millisecond * 10)
			instances[index] = singleton.GetInstance()
		}(i)
	}

	wg.Wait()

	// Verify all goroutines got the same instance
	allSame := true
	firstID := instances[0].GetConnectionID()
	for i, inst := range instances {
		if inst.GetConnectionID() != firstID {
			allSame = false
			fmt.Printf("   ✗ Instance %d has different ID: %d\n", i, inst.GetConnectionID())
		}
	}

	if allSame {
		fmt.Printf("   ✓ All %d goroutines received the same instance (ID: %d)\n", len(instances), firstID)
		fmt.Println("   ✓ Thread-safety verified!\n")
	}

	// Demonstrate that the instance persists across function calls
	fmt.Println("4. Instance persistence:")
	fmt.Printf("   Connection string: %s\n", db1.GetConnectionString())
	fmt.Printf("   Connection ID: %d\n", db1.GetConnectionID())
	fmt.Println("   ✓ The same instance is reused across the entire program lifecycle")
}
