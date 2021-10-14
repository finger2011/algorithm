package week03onclass

// https://leetcode-cn.com/problems/course-schedule/
// leetcode 207. 课程表

func canFinish(numCourses int, prerequisites [][]int) bool {
	var toCoureses = make(map[int][]int, numCourses)
	var inDegOfCoureses = make(map[int]int, numCourses)
	for i := 0; i < numCourses; i++ {
		inDegOfCoureses[i] = 0
	}
	for _, pre := range prerequisites {
		if _, ok := toCoureses[pre[0]]; ok {
			toCoureses[pre[0]] = append(toCoureses[pre[0]], pre[1])
		} else {
			toCoureses[pre[0]] = []int{pre[1]}
		}
		inDegOfCoureses[pre[1]]++
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
	return len(lesson) == numCourses
}
