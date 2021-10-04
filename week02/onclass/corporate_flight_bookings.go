package week02onclass

// https://leetcode-cn.com/problems/corporate-flight-bookings/
// leetcode 1109 航班预订统计

func corpFlightBookings(bookings [][]int, n int) []int {
	var s = make([]int, n+1)
	for i := 0; i < len(bookings); i++ {
		s[bookings[i][0]] += bookings[i][2]
		if bookings[i][1] < n {
			s[bookings[i][1]+1] -= bookings[i][2]
		}
	}
	for i := 1; i < len(s); i++ {
		s[i] = s[i] + s[i-1]
	}
	s = s[1:]
	return s
}
