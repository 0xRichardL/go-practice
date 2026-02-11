package main

// Object represents a simple structure to allocate
type Object struct {
	Data []byte
	Next *Object
}

// RapidSmallAllocations creates many small objects quickly
func RapidSmallAllocations() []*Object {
	objects := make([]*Object, 100000)
	for i := 0; i < 100000; i++ {
		objects[i] = &Object{
			Data: make([]byte, 100),
		}
	}
	return objects
}
