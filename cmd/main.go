package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/DataDog/datadog-pgo/internal"
)

// main runs the pgo tool.
func main() {
	if err := internal.RunMainCmdLine(); err != nil && !errors.As(err, &internal.HandledError{}) {
		if !errors.As(err, &internal.LoggedError{}) {
			fmt.Fprintf(os.Stderr, "pgo: error: %v\n", err)
		}
		os.Exit(1)
	}
}
