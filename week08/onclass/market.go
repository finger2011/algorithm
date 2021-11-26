package week08onclass

// https://www.acwing.com/problem/content/147/
// acwing 145. 超市
var marketFa []int

func market(products [][]int) int {
	marketFa = make([]int, 10001)
	for i := 0; i <= 10000; i++ {
		marketFa[i] = i
	}
	var ans int
	quickSort(products, 0, len(products) - 1)
	for i := len(products) - 1; i >= 0; i-- {
		profit, day := products[i][0], products[i][1]
		lastDay := marketFind(day)
		if lastDay > 0 {
			ans += profit
			marketFa[lastDay] = lastDay - 1
		}
	}
	return ans
}

func quickSort(values [][]int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j][0] >= temp[0] {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		for i <= p && values[i][0] <= temp[0] {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func marketFind(x int) int {
	if marketFa[x] != x {
		marketFa[x] = marketFind(marketFa[x])
	}
	return marketFa[x]
}