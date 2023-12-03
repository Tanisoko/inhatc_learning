package main

import "fmt"

func main() {
	// var primes [3]int
	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5
	// fmt.Println(primes, primes[1])

	primes := [3]int{2, 3, 5}
	fmt.Println(primes, primes[1])

	test := [5]bool{true, true, true}
	fmt.Println(test[3])
	// fmt.Println(test)

	i := 0
	for i < len(primes) {
		fmt.Println(primes[i])
		i++
	}

	for j := 0; j < len(primes); j++ {
		fmt.Println(primes[j])
	}

	// for idx, prime := range primes {
	// 	fmt.Println(idx, prime)
	// }

	for _, prime := range primes {
		fmt.Println(prime)
	}

	fmt.Printf("%#v\n", test)
	fmt.Printf("%#v\n", primes)

}
