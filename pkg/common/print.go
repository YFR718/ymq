package common

import (
	"fmt"
	"runtime"
)

func PrintError(err error) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("An error occurred in %s:%d: %v\n", file, line, err)
}
