package week05

// https://leetcode-cn.com/problems/online-election/
// leetcode 911. 在线选举

//TopVotedCandidate TopVotedCandidate
type TopVotedCandidate struct {
	persons []int
	times   []int
	result  []int
}

//Constructor Constructor
func Constructor(persons []int, times []int) TopVotedCandidate {
	var obj = TopVotedCandidate{persons: persons, times: times}
	var tickets = make(map[int]int, 3)
	var maxPerson = persons[0]
	obj.result = append(obj.result, maxPerson)
	tickets[maxPerson]++
	for i := 1; i < len(persons); i++ {
		tickets[persons[i]]++
		if persons[i] == maxPerson {
			obj.result = append(obj.result, maxPerson)
		} else if tickets[persons[i]] >= tickets[maxPerson] {
			obj.result = append(obj.result, persons[i])
			maxPerson = persons[i]
		} else {
			obj.result = append(obj.result, maxPerson)
		}
	}
	return obj
}

//Q q
func (top *TopVotedCandidate) Q(t int) int {
	var left, right = 0, len(top.result) - 1
	for left < right {
		var mid = (left + right + 1) / 2
		if top.times[mid] == t {
			return top.result[mid]
		} else if top.times[mid] > t {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return top.result[right]
}
