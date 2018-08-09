package main

func numDiv(start, finish, divisor int) int {
	plusNumber := 0
	if start%divisor == 0 {
		plusNumber = 1
	}
	return finish/divisor - start/divisor + plusNumber
}
