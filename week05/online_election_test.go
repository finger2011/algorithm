package week05

import (
	"fmt"
	"testing"
)

func TestTopVotedCandidate_Q(t *testing.T) {
	persons := []int{0, 1, 1, 0, 0, 1, 0}
	times := []int{0, 5, 10, 15, 20, 25, 30}
	obj := Constructor(persons, times)
	params := []int{3, 12, 25, 15, 24, 8}
	result := []int{0, 1, 1, 0, 0, 1}
	for i := 0; i < len(params); i++ {
		if obj.Q(params[i]) != result[i] {
			t.Errorf("param:%d, want:%d, got:%d\n", params[i], result[i], obj.Q(params[i]))
		}
	}
}

func TestTopVotedCandidate_Q1(t *testing.T) {
	persons := []int{0, 0, 1, 1, 2}
	times := []int{0, 67, 69, 74, 87}
	obj := Constructor(persons, times)
	fmt.Printf("result:%v\n", obj.result)
	params := []int{4, 62, 100, 88, 70, 73, 22, 75, 29, 10}
	result := []int{0, 0, 1, 1, 0, 0, 0, 1, 0, 0}
	for i := 0; i < len(params); i++ {
		if obj.Q(params[i]) != result[i] {
			t.Errorf("param:%d, want:%d, got:%d\n", params[i], result[i], obj.Q(params[i]))
		}
	}
}
