package main

import "fmt"

type Student struct {
	Name   string
	Points []int
}

func main() {
	studentA := Student{
		Name:   "Student A",
		Points: []int{3, 5, 1},
	}
	studentB := studentA
	studentA.Points[2] = 10

	fmt.Println(studentB)

}

// NOTE:
// Stubborn elder interviewer, only focus on language's unreal edge-cases.
// No algorithms, no system design, no new tech.
// Prefer old zombies rather than new bloods.
// 1. Goroutine edge-cases
// 2. Slice edge-cases
// 3. DDD concept: Aggregation root, value object
