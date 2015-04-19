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

func (fizzbuzz *FizzBuzz) FromString(s string) string {
	i, _ := strconv.Atoi(s)

	return fizzbuzz.Parse(i)
}
