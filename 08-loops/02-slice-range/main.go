package main

func Sum(digit ...int) int {
	var sum int = 0

	for _, d := range digit {
		sum += d
	}

	return sum
}

func main() {
	_ = Sum(1, 2, 3, 4, 5)
}
