package main

import (
	"fmt"
	"os"
)

func main() {
	fizzbuzz := FizzBuzz{}

	fmt.Println(fizzbuzz.FromString(os.Args[1]))
}
