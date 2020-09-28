package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var aux, factorial int
	factorial = 1
	aux = n
	var sum float64
	sum = 0
	for i := 0; i <= n; i++ {
		for j := aux; j > 0; j-- {
			factorial *= j
		}
		sum += float64(1) / float64(factorial)
		aux--
		factorial = 1
	}
	fmt.Printf(" %f", sum)

}
