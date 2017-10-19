package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	defer file.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
