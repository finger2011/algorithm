package week03

// https://leetcode-cn.com/problems/course-schedule-ii/
// leetcode 210. 课程表 II

func findOrder(numCourses int, prerequisites [][]int) []int {
	var toCoureses = make(map[int][]int, numCourses)
	var inDegOfCoureses = make(map[int]int, numCourses)
	for i := 0; i < numCourses; i++ {
		inDegOfCoureses[i] = 0
	}
	for _, pre := range prerequisites {
		if _, ok := toCoureses[pre[1]]; ok {
			toCoureses[pre[1]] = append(toCoureses[pre[1]], pre[0])
		} else {
			toCoureses[pre[1]] = []int{pre[0]}
		}
		inDegOfCoureses[pre[0]]++
	}
	var queue []int
	for i := 0; i < len(inDegOfCoureses); i++ {
		if inDegOfCoureses[i] == 0 {
			queue = append(queue, i)
		}
	}
	var lesson []int
	for len(queue) > 0 {
		var x = queue[0]
		queue = queue[1:]
		lesson = append(lesson, x)
		for _, y := range toCoureses[x] {
			inDegOfCoureses[y]--
			if inDegOfCoureses[y] == 0 {
				queue = append(queue, y)
			}
		}
	}
	if len(lesson) == numCourses {
		return lesson
	}
	return []int{}
}
