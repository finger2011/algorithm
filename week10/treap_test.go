package week10

import (
	"testing"
)

func Test_newTreap(t *testing.T) {
	tree := newTreap()
	tree = tree.insert(10)
	tree = tree.insert(20)
	tree = tree.insert(30)
	var num int 
	num = tree.getRankByVal(20)
	if num != 2 {
		t.Errorf("getRankByVal() = %d, want %d", num, 2)
	}
	num, _ = tree.getValByRank(3)
	if num != 20 {
		t.Errorf("getValByRank() = %d, want %d", num, 20)
	}

}
