package main

import (
	"fmt"
	"os"
)

// ExitOnError will exit the application if err is not nil
func ExitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
