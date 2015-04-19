package main

import "testing"

func TestParse1(t *testing.T) {
	fizzbuzz := FizzBuzz{}

	result := fizzbuzz.Parse(1)

	if result != "1" {
		t.Error("FizzBuzz parse error", "1", result)
	}
}

func TestParse2(t *testing.T) {
	fizzbuzz := FizzBuzz{}

	result := fizzbuzz.Parse(2)

	if result != "2" {
		t.Error("FizzBuzz parse error", "2", result)
	}
}

func TestParse3(t *testing.T) {
	fizzbuzz := FizzBuzz{}

	result := fizzbuzz.Parse(3)

	if result != "Fizz" {
		t.Error("FizzBuzz parse error", "Fizz", result)
	}
}

func TestParse5(t *testing.T) {
	fizzbuzz := FizzBuzz{}

	result := fizzbuzz.Parse(5)

	if result != "Buzz" {
		t.Error("FizzBuzz parse error", "Buzz", result)
	}
}

func TestParse15(t *testing.T) {
	fizzbuzz := FizzBuzz{}

	result := fizzbuzz.Parse(15)

	if result != "FizzBuzz" {
		t.Error("FizzBuzz parse error", "FizzBuzz", result)
	}
}

func TestSequence(t *testing.T) {
	fizzbuzz := FizzBuzz{}

	result := fizzbuzz.Sequence(1, 5)

	if result != "1 2 Fizz 4 Buzz" {
		t.Error("FizzBuzz sequence error", "1 2 Fizz 4 Buzz", result)
	}
}
