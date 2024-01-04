package fizz_buzz

import (
	"strconv"
)

func intToString2(n int) string {
	return strconv.Itoa(n)
}

func fizzBuzz2(n int) []string {
	r := make([]string, n)

	fn := func(i int) string {
		if i%15 == 0 {
			return "FizzBuzz"
		}

		if i%5 == 0 {
			return "Buzz"
		}

		if i%3 == 0 {
			return "Fizz"
		}

		return strconv.Itoa(i)
	}

	for i := 1; i <= n; i++ {
		r[i-1] = fn(i)
	}

	return r
}

func fizzBuzz(n int) []string {
	r := make([]string, 0, n)

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			r = append(r, "FizzBuzz")
		} else if i%5 == 0 {
			r = append(r, "Buzz")
		} else if i%3 == 0 {
			r = append(r, "Fizz")
		} else {
			r = append(r, intToString2(i))
		}
	}

	return r
}
