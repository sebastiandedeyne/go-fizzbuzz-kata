package main

import (
	"strconv"
)

type FizzBuzz struct {
}

func (_ *FizzBuzz) Parse(i int) string {
	if i%15 == 0 {
		return "FizzBuzz"
	}

	if i%3 == 0 {
		return "Fizz"
	}

	if i%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(i)
}

func (fb *FizzBuzz) Sequence(start int, end int) string {
	sequence := fb.Parse(start)

	for i := start+1; i <= end; i++ {
		sequence = sequence + " " + fb.Parse(i)
	}

	return sequence
}
