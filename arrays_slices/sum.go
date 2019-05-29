package arrays_slices

// Sum takes any size of slice in integer and returns the sum of them
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll takes a varying number of slices, and returns a new slice containing
// the totals for each slice passed in
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

// SumAllTails takes a varying number of slices, and returns a new slice containing
// the totals of the "tails" of each slice passed in. The tail of a collection is
// all the items apart from the first one(the "head")
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
