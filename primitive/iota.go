package main

// Start from 0 and increment by 1
const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
)

// Start from 1 and shift left by 1 each time
const (
	Read    = 1 << iota // 1
	Write               // 2
	Execute             // 4
)

// Start from 1 and increment by 1
const (
	First  = iota + 1 // 1
	Second            // 2
	Third             // 3
)

// Note: Each group of constants/variables resets iota.
