package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	model := bson.M{
		"field_1": 50,
		"field_2": nil,
	}
	model["field_3"] = nil
	fmt.Println(model)
}
