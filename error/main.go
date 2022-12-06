package main

import (
	"errors"
	"fmt"
)

func UnwrapTillCore(err error) error {
	if inner := errors.Unwrap(err); inner != nil {
		return UnwrapTillCore(inner)
	}
	return err
}

func main() {
	err := errors.New("CORE")
	err = fmt.Errorf("L1: %w", err)
	err = fmt.Errorf("L2: %w", err)
	err = fmt.Errorf("L3: %w", err)
	err = fmt.Errorf("L4: %w", err)
	fmt.Println(UnwrapTillCore(err))
}
