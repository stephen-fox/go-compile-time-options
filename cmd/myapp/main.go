package main

import (
	"fmt"

	"github.com/stephen-fox/myapp/cmd/myapp/compileoptions"
)

func main() {
	fmt.Printf("DisableMTLS is %v\n", compileoptions.DisableMTLS())
}
