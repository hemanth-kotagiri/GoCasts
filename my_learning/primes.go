package main

import (
	"fmt"
	"math"
)

func checkPrime(n int) bool {
	for i := 2; i < int(math.Pow(float64(n), 0.5))+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func genPrimes(n int) []int {
	primes := []int{}
	for i := 1; i <= n; i++ {
		if checkPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	fmt.Println(genPrimes(30))
}
