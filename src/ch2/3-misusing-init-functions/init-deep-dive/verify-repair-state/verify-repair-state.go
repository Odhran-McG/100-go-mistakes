// verify_repair_state.go
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// =========================================================================
// Use Case 2: Verify or Repair Program State
// =========================================================================
// The init() function is commonly used to verify the program's state and
// repair it if necessary before execution begins. This includes:
// - Checking required environment variables
// - Setting default values
// - Verifying system integrity
// - Ensuring required resources exist

// Example 2.1: Environment configuration validation and defaults
// This example demonstrates how to check for required environment variables
// and provide sensible defaults when possible.

// Package-level variables initialized from environment
var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

// First init function - validates environment configuration
// This demonstrates:
// - Checking for required variables (USER)
// - Providing default values when possible (HOME, GOPATH)
// - Setting up command-line flags to override defaults
func init() {
	fmt.Println("Running init for environment validation...")
	
	// Check if required variables are set
	if user == "" {
		// Fatal error for missing critical variable
		// In a real application, you might log this and exit
		log.Println("$USER not set - would normally exit")
		user = "defaultuser" // Using default for demonstration
	}
	
	// Set default values for optional variables
	if home == "" {
		home = "/home/" + user
		fmt.Printf("HOME not set, using default: %s\n", home)
	}
	
	if gopath == "" {
		gopath = home + "/go"
		fmt.Printf("GOPATH not set, using default: %s\n", gopath)
	}
	
	// Set up command-line flag to override gopath
	// This shows how init can be used to configure the program
	flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}

// Example 2.2: Component registration and system integrity check
// This example shows how to verify that all required components
// are registered and necessary resources exist before execution.

// Placeholder handler types for demonstration
type UserHandler struct{}
type ProductHandler struct{}
type OrderHandler struct{}
type AuthHandler struct{}

func (h UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
func (h ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
func (h OrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
func (h AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

// Package-level map to hold registered handlers
var registeredHandlers map[string]http.Handler

// Second init function - registers components and verifies integrity
// This demonstrates:
// - Setting up a component registry
// - Verifying all required components exist
// - Ensuring required resources (directories) exist
func init() {
	fmt.Println("Running init for system integrity check...")
	
	// Initialize and register API handlers
	registeredHandlers = make(map[string]http.Handler)
	registeredHandlers["/api/users"] = UserHandler{}
	registeredHandlers["/api/products"] = ProductHandler{}
	registeredHandlers["/api/orders"] = OrderHandler{}
	registeredHandlers["/api/auth"] = AuthHandler{}

	// Verify all required handlers are registered
	// This is a critical integrity check before the program runs
	requiredEndpoints := []string{"/api/users", "/api/products", "/api/orders", "/api/auth"}
	for _, endpoint := range requiredEndpoints {
		if _, exists := registeredHandlers[endpoint]; !exists {
			// In a real application, this would be a fatal error
			log.Printf("Missing required handler for endpoint: %s\n", endpoint)
		}
	}

	// Verify filesystem resources exist, create if needed
	// This repairs the program state by ensuring required directories exist
	dirs := []string{"./uploads", "./logs", "./temp"}
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Printf("Creating required directory: %s\n", dir)
			if err := os.MkdirAll(dir, 0755); err != nil {
				// In a real application, this would be a fatal error
				log.Printf("Failed to create directory %s: %v\n", dir, err)
			}
		}
	}
}

// Display the verified/repaired state when imported
func DisplayStateData() {
	fmt.Println("\n=== Verify/Repair State Results ===")
	
	fmt.Println("\nConfiguration values:")
	fmt.Printf("User: %s\n", user)
	fmt.Printf("Home: %s\n", home)
	fmt.Printf("Gopath: %s\n", gopath)
	
	fmt.Println("\nRegistered API endpoints:")
	for endpoint := range registeredHandlers {
		fmt.Printf("- %s\n", endpoint)
	}
	
	fmt.Println("\nVerified directories:")
	dirs := []string{"./uploads", "./logs", "./temp"}
	for _, dir := range dirs {
		if _, err := os.Stat(dir); err == nil {
			fmt.Printf("- %s (exists)\n", dir)
		} else {
			fmt.Printf("- %s (missing)\n", dir)
		}
	}
}

// This main function only executes when running this file directly
func main() {
	fmt.Println("Verify/Repair Program State Examples")
	flag.Parse()
	DisplayStateData()
}
