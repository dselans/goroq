package main

import (
	"fmt"
)

const (
	VERSION string = "0.0.1"
)

func main() {
	opts := handleCliArgs()
	fmt.Println(opts)
}
