package main

func Sum(numbers []int) int {
	sum := 0
	//On each iteration, range returns two values - the index and the value. We are choosing to ignore the index value by using _ blank identifier.
	for _, number := range numbers { //range lets you iterate over an array.
		sum += number
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
