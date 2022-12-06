package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type A struct {
	A string `json:"a,omitempty"`
}

func main() {
	base := `{"a":"b"}`
	bStream := strings.NewReader(base)
	sStream := new(bytes.Buffer)
	io.Copy(sStream, bStream)
	fmt.Println("sStream:", sStream.String())
	a := &A{}
	if err := json.NewDecoder(bStream).Decode(a); err != nil {
		fmt.Println("json.NewDecoder(bStream).Decode: %w", err)
	}
	fmt.Println("a", a)

}

type Constraint interface {
	int | ~string
}

func test[A Constraint](input1, input2 A) bool {
	return input1 > input2
}
