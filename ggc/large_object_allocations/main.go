package main

// Object represents a simple structure to allocate
type Object struct {
	Data []byte
	Next *Object
}

// LargeObjectAllocations creates fewer but larger objects
func LargeObjectAllocations() []*Object {
	objects := make([]*Object, 100)
	for i := 0; i < 100; i++ {
		objects[i] = &Object{
			Data: make([]byte, 1024*1024), // 1 MB each
		}
	}
	return objects
}
