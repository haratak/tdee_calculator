package main

import (
	"fmt"
	"os"
)

const (
	Male = iota
	Female
)

func main() {
	args := os.Args

	for i, arg := range args {
		fmt.Println(i, arg)
	}
}
