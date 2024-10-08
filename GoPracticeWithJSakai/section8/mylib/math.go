package mylib

func Average(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total / len(x)
}
