package slices

// Sum returns the sum of all the integers in array numbers.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
