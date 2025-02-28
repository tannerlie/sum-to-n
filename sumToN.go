package main

import "fmt"

/*
This method uses a niave method with a for loop that increments through every number from 1 to n and adds that number to the variable "sum".
It runs in O(n) time.
*/
func sum_to_n_a(n int) int { // Incremental for loop method
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

/*
This method uses mathematical calculation to pair up the numbers on either ends of the sequence. Essentially there are n/2 pairs summing to (n + 1)
It runs in O(1) time.
*/
func sum_to_n_b(n int) int {
	return n * (n + 1) / 2
}

/*
This method uses a divide-and-conquer method using recursion. It seperates the sequence into 2 halfs to calculate seperately however it still visits every number at least once.
The time complexity is still O(n)
*/
func sum_to_n_c(n int) int {

	var sum_of_sequence func(int, int) int
	sum_of_sequence = func(start int, end int) int {
		if start == end {
			return start
		}
		mid := start + (end-start)/2
		return sum_of_sequence(start, mid) + sum_of_sequence(mid+1, end)
	}

	return sum_of_sequence(1, n)
}

func main() {
	var n = 112
	fmt.Println(sum_to_n_a(n))
	fmt.Println(sum_to_n_b(n))
	fmt.Println(sum_to_n_c(n))
}
