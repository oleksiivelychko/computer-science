package factorial

func Factorial(x uint) uint {
	if x == 1 {
		return 1
	}
	return x + Factorial(x-1)
}