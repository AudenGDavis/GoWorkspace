package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := "Password123"

	stringHash := fmt.Sprintf("%x\n", sha256.Sum256([]byte(data)))

	print(stringHash)

}
