package main

import (
	"fmt"
	"os"
)

func main() {
	usefulArgs := os.Args[1:]
	fmt.Println(len(usefulArgs), usefulArgs)
}
