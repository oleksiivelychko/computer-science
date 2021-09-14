package factorial

func Factorial(x uint8) uint8 {
	if x == 1 {
		return 1
	}
	return x + Factorial(x-1)
}
