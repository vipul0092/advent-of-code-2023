package utils

func LCM(numbers []int) int {
	current := numbers[0]
	for i := 1; i < len(numbers); i++ {
		current = (current * numbers[i]) / gcd(current, numbers[i])
	}
	return current
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
