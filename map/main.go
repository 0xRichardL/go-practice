package main

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrCredentialInvalid  = errors.New("a")
	ErrDuplicatedUsername = errors.New("b")
	ErrLoginFailed        = errors.New("c")
)

var errorStatusMap = map[error]int{
	ErrCredentialInvalid:  http.StatusBadRequest,
	ErrDuplicatedUsername: http.StatusBadRequest,
	ErrLoginFailed:        http.StatusServiceUnavailable,
}

func main() {
	fmt.Println(errorStatusMap[nil])
}
