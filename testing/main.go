package main

func Sum(x, y int) int {
	return x + y
}

func GetMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Fibonnaci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonnaci(n-1) + Fibonnaci(n-2)

}