package main

import (
	"fmt"

	"github.com/mp40/dkk-to-aud/read"
)

func main() {
	values := read.Read()

	fmt.Println(values)
}
