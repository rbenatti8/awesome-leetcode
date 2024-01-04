package richest_customer_wealth

func sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}

	return s
}

func sum2(nums []int, c chan int) {
	s := 0
	for _, num := range nums {
		s += num
	}

	c <- s
}

func maximumWealth(accounts [][]int) int {
	maxWealth := 0

	for _, account := range accounts {
		s := sum(account)
		if s > maxWealth {
			maxWealth = s
		}
	}

	return maxWealth
}

func maximumWealth2(accounts [][]int) int {
	maxWealth := 0

	for _, account := range accounts {
		c := make(chan int)

		go sum2(account, c)

		s := <-c

		if s > maxWealth {
			maxWealth = s
		}

		close(c)

	}

	return maxWealth
}
