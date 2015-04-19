package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fb := FizzBuzz{}

	if len(os.Args) == 2 {
		i, _ := strconv.Atoi(os.Args[1])
		fmt.Println(fb.Parse(i))
		return
	}

	if len(os.Args) == 3 {
		start, _ := strconv.Atoi(os.Args[1])
		end, _ := strconv.Atoi(os.Args[2])
		fmt.Println(fb.Sequence(start, end))
		return
	}
}
