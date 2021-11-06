package main

import (
	"fmt"

	"os"

	"github.com/djworth/riddler/pkg/hash"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: hasher <string_to_hash>")
		return
	}

	fmt.Println(hash.Hash(os.Args[1]))
}
