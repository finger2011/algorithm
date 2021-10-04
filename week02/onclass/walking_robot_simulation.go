package week02onclass

import (
	"strconv"
)

//https://leetcode-cn.com/problems/walking-robot-simulation/
// leetcode 874 模拟行走机器人

func robotSim(commands []int, obstacles [][]int) int {
	var ant int
	var obs = make(map[string]int, len(obstacles))
	for i := 0; i < len(obstacles); i++ {
		obs[strconv.Itoa(obstacles[i][0])+","+strconv.Itoa(obstacles[i][1])] = 1
	}
	var x, y int
	var dir int // N = 0 E = 1 S = 2 W = 3
	var dx = []int{0, 1, 0, -1}
	var dy = []int{1, 0, -1, 0}
	for i := 0; i < len(commands); i++ {
		if -1 == commands[i] {
			dir = (dir + 1) % 4
		} else if -2 == commands[i] {
			dir = (dir + 3) % 4
		} else {
			for j := 0; j < commands[i]; j++ {
				var nx = x + dx[dir]
				var ny = y + dy[dir]
				if _, ok := obs[strconv.Itoa(nx)+","+strconv.Itoa(ny)]; ok {
					break
				}
				x = nx
				y = ny
			}
		}
		if x*x+y*y > ant {
			ant = x*x + y*y
		}
	}
	return ant
}

func robotSim2(commands []int, obstacles [][]int) int {
	var ant int
	var obs = make(map[int64]int, len(obstacles))
	for i := 0; i < len(obstacles); i++ {
		obs[calHash(obstacles[i])] = 1
	}
	var x, y int
	var dir int // N = 0 E = 1 S = 2 W = 3
	var dx = []int{0, 1, 0, -1}
	var dy = []int{1, 0, -1, 0}
	for i := 0; i < len(commands); i++ {
		if -1 == commands[i] {
			dir = (dir + 1) % 4
		} else if -2 == commands[i] {
			dir = (dir + 3) % 4
		} else {
			for j := 0; j < commands[i]; j++ {
				var nx = x + dx[dir]
				var ny = y + dy[dir]
				if _, ok := obs[calHash([]int{nx, ny})]; ok {
					break
				}
				x = nx
				y = ny
			}
		}
		if x*x+y*y > ant {
			ant = x*x + y*y
		}
	}
	return ant
}

func calHash(obj []int) int64 {
	return int64((obj[0]+30000)*60001 + (obj[1] + 30000))
}
