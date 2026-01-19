package singleton

import (
	"fmt"
	"sync"
)

// DatabaseConnection represents a singleton database connection
type DatabaseConnection struct {
	connectionString string
	isConnected      bool
	connectionID     int
}

var (
	instance *DatabaseConnection
	once     sync.Once
	connID   int
)

// GetInstance returns the singleton instance of DatabaseConnection
// This is thread-safe and will only create the instance once
func GetInstance() *DatabaseConnection {
	once.Do(func() {
		connID++
		instance = &DatabaseConnection{
			connectionString: "postgresql://localhost:5432/mydb",
			isConnected:      false,
			connectionID:     connID,
		}
		fmt.Printf("Database connection instance created (ID: %d)\n", connID)
	})
	return instance
}

// Connect simulates connecting to the database
func (db *DatabaseConnection) Connect() {
	if !db.isConnected {
		db.isConnected = true
		fmt.Printf("Connected to database (ID: %d)\n", db.connectionID)
	} else {
		fmt.Printf("Already connected to database (ID: %d)\n", db.connectionID)
	}
}

// Disconnect simulates disconnecting from the database
func (db *DatabaseConnection) Disconnect() {
	if db.isConnected {
		db.isConnected = false
		fmt.Printf("Disconnected from database (ID: %d)\n", db.connectionID)
	}
}

// Query simulates executing a database query
func (db *DatabaseConnection) Query(sql string) {
	if !db.isConnected {
		fmt.Println("Error: Not connected to database. Call Connect() first.")
		return
	}
	fmt.Printf("Executing query: %s (Connection ID: %d)\n", sql, db.connectionID)
}

// GetConnectionID returns the unique connection ID
func (db *DatabaseConnection) GetConnectionID() int {
	return db.connectionID
}

// GetConnectionString returns the connection string
func (db *DatabaseConnection) GetConnectionString() string {
	return db.connectionString
}
