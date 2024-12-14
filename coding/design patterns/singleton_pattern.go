package main

import (
	"fmt"
	"sync"
)

// Store represents a key-value store
type Store struct {
	data map[string]string
	mu   sync.RWMutex // Mutex to handle concurrent access
}

// Singleton instance and sync.Once for thread-safe initialization
var storeInstance *Store
var once sync.Once

// GetStoreInstance provides a thread-safe global access to the store instance
func GetStoreInstance() *Store {
	once.Do(func() {
		// Initialize the store instance
		storeInstance = &Store{
			data: make(map[string]string),
		}
	})
	return storeInstance
}

// Set sets a key-value pair in the store
func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

// Get retrieves a value by key from the store
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, exists := s.data[key]
	return value, exists
}

// Main function to test the Singleton Store
func main() {
	// Get the single store instance
	store1 := GetStoreInstance()
	store1.Set("username", "john_doe")

	// Get the instance again
	store2 := GetStoreInstance()
	value, exists := store2.Get("username")

	if exists {
		fmt.Printf("Key 'username' has value: %s\n", value)
	} else {
		fmt.Println("Key 'username' does not exist.")
	}

	// Confirm both store1 and store2 are the same instance
	if store1 == store2 {
		fmt.Println("Both store instances are the same.")
	} else {
		fmt.Println("Store instances are different.")
	}
}
