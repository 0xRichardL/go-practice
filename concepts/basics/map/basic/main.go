package main

import "fmt"

// Underlying structure of a Go map (simplified version):
// A Go map is a hash table, implemented by the runtime as:
//
//	type hmap struct {
//	    count     int        // number of key-value pairs
//	    flags     uint8
//	    B         uint8      // log2(number of buckets)
//	    noverflow uint16
//	    hash0     uint32     // random hash seed
//	    buckets    *bmap     // array of buckets
//	    oldbuckets *bmap     // during growth
//	    nevacuate  uintptr
//	}
//
// Each bucket (bmap) holds:
// * 8 key/value pairs
// * 8 tophash bytes (1 byte per slot)
// * Keys stored together
// * Values stored together
// * Overflow pointer if full
func main() {
	// Create a map
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30

	// Map literal
	scores := map[string]int{
		"math":    95,
		"english": 87,
	}

	// Access values
	fmt.Println("Alice's age:", ages["Alice"])
	fmt.Println("Math score:", scores["math"])

	// Check if key ok
	age, ok := ages["Charlie"]
	if ok {
		fmt.Println("Charlie's age:", age)
	} else {
		fmt.Println("Charlie not found")
	}

	// Update value
	ages["Alice"] = 26

	// Delete key
	delete(ages, "Bob")

	// Iterate over map.
	// Will receive random order due to map buckets migration during growth (rehashing).
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// Get map length
	fmt.Println("Number of people:", len(ages))
}
