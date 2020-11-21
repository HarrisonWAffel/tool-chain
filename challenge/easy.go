package challenge

import (
	"fmt"
)

func FizzBuzz(n int) []interface{} {
	var Buzzed []interface{}
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			Buzzed = append(Buzzed, "FizzBuzz")
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			Buzzed = append(Buzzed, "Buzz")
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			Buzzed = append(Buzzed, "Fizz")
			fmt.Println("Fizz")
		} else {
			Buzzed = append(Buzzed, i)
			fmt.Println(i)
		}
	}
	return Buzzed
}

func Primes(n int) []int {
	var primes []int
	for i := 1; i <= n; i++ {
		var prime = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				prime = false
			}
		}
		if prime {
			primes = append(primes, i)
		}
	}
	return primes
}
