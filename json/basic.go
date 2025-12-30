package main

import (
	"encoding/json"
	"fmt"
)

// Comprehensive struct demonstrating all JSON serialization use cases
type User struct {
	// Basic JSON tags
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`

	// Optional fields with omitempty
	Age         int      `json:"age,omitempty"`
	PhoneNumber string   `json:"phone_number,omitempty"`
	Tags        []string `json:"tags,omitempty"`      // Omit if empty slice
	IsActive    *bool    `json:"is_active,omitempty"` // Omit if nil

	// String option - convert number to/from string in JSON
	Balance float64 `json:"balance,string"`

	// Nested struct (2-layer structure)
	Address Address `json:"address"`

	// Ignored field - won't appear in JSON
	Password string `json:"-"`
}

// Nested struct (second layer)
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
	ZipCode string `json:"zip_code,omitempty"` // Optional field
}

func main() {
	// Example 1: Full struct with all fields populated
	isActive := true
	user1 := User{
		ID:          1,
		Username:    "john_doe",
		Email:       "john@example.com",
		Age:         30,
		PhoneNumber: "+1-555-0123",
		Tags:        []string{"premium", "verified"},
		IsActive:    &isActive,
		Balance:     1250.75,
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
			ZipCode: "10001",
		},
		Password: "secret123", // Won't appear in JSON (json:"-")
	}

	// Marshal: Go struct -> JSON
	jsonData, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}
	fmt.Println("1. Compact JSON:", string(jsonData))

	// Marshal with indentation (pretty print)
	jsonIndented, _ := json.MarshalIndent(user1, "", "  ")
	fmt.Println("\n2. Indented JSON (all fields):\n", string(jsonIndented))
	fmt.Println("   Note: Password field is excluded (json:\"-\" tag)")
	fmt.Println("   Note: Balance is a string (string tag option)")

	// Example 2: Struct with omitempty fields not set
	user2 := User{
		ID:       2,
		Username: "jane_doe",
		Email:    "jane@example.com",
		Balance:  500.00,
		Address: Address{
			Street:  "456 Oak Ave",
			City:    "Los Angeles",
			Country: "USA",
			// ZipCode omitted - won't appear due to omitempty
		},
		// Age, PhoneNumber, Tags, IsActive omitted - won't appear due to omitempty
	}

	data2, _ := json.MarshalIndent(user2, "", "  ")
	fmt.Println("\n3. JSON with omitempty fields excluded:\n", string(data2))
	fmt.Println("   Note: age, phone_number, tags, is_active, zip_code excluded")

	// Unmarshal: JSON -> Go struct
	jsonStr := `{
		"id": 3,
		"username": "bob_smith",
		"email": "bob@example.com",
		"balance": "750.25",
		"address": {
			"street": "789 Pine Rd",
			"city": "Chicago",
			"country": "USA",
			"zip_code": "60601"
		},
		"tags": ["new", "standard"]
	}`

	var user3 User
	err = json.Unmarshal([]byte(jsonStr), &user3)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}
	fmt.Println("\n4. Unmarshaled struct:")
	fmt.Printf("   User: %+v\n", user3)
	fmt.Printf("   Address: %+v\n", user3.Address)
	fmt.Printf("   Balance (parsed from string): %.2f\n", user3.Balance)
}
