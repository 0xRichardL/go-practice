package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Data struct {
	A string `json:"a,omitempty"`
}

func main() {
	// Define a base JSON string.
	base := `{"a":"b"}`

	// Create a new strings.Reader from the base string.
	// A strings.Reader implements the io.Reader interface.
	aStream := strings.NewReader(base)

	// Create a new bytes.Buffer. This acts as an in-memory buffer that can be written to and read from.
	bStream := new(bytes.Buffer)

	// Copy the contents from aStream (strings.Reader) to bStream (bytes.Buffer).
	// After this operation, aStream's internal pointer will be at the end,
	// meaning it has been fully read.
	io.Copy(bStream, aStream) // aStream is now at EOF

	// Create a new instance of struct A.
	a := &Data{}

	// Attempt to decode JSON from aStream into struct a.
	// Since aStream was fully read by io.Copy, it's at EOF.
	// Decoding from an empty reader will result in an EOF error.
	if err := json.NewDecoder(aStream).Decode(a); err != nil { // This will return an EOF error.
		fmt.Printf("Decode aStream: %v\n", err)
	}
	// Print the struct a. It will likely be empty or its default values
	// because the decoding failed.
	fmt.Println("a data:", a) // 'a' will be an empty struct due to the decoding error.

	// Print the content of bStream. It should contain "{"a":"b"}".
	fmt.Println("bStream:", bStream.String())

	b := &Data{}
	// Decode JSON from bStream into struct b.
	// bStream contains the original JSON string, so this decoding should succeed.
	if err := json.NewDecoder(bStream).Decode(b); err != nil {
		fmt.Printf("Decode bStream: %v\n", err)
	}

	fmt.Println("b data:", b) // 'b' will contain the decoded data: {A:b}.
}
