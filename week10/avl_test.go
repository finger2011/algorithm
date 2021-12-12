package week10

import (
	"testing"
)

func Test_newAvlTree3(t *testing.T) {
	var num int
	tree := newAvlTree()
	tree = tree.insert(577793)
	tree = tree.insert(408221)
	tree = tree.insert(880861)
	tree = tree.delete(408221)
	tree = tree.insert(460353)
	tree = tree.insert(223489)
	num = tree.getValByRank(2)
	if num != 460353 {
		t.Errorf("getValByRank() = %d, want %d", num, 460353)
	}
	// tree = tree.insert(460353)
	// tree = tree.insert(460353)
	// tree = tree.insert(460353)
	// tree = tree.insert(460353)
	// tree = tree.insert(460353)
	// tree = tree.insert(460353)
	
	// num = tree.getRankByVal(964673)
	// if num != 3 {
	// 	t.Errorf("getRankByVal() = %d, want %d", num, 3)
	// }
}

func Test_newAvlTree2(t *testing.T) {
	tree := newAvlTree()
	tree = tree.insert(964673)

	tree = tree.insert(915269)
	tree = tree.insert(53283)
	var num int
	num = tree.getRankByVal(964673)
	if num != 3 {
		t.Errorf("getRankByVal() = %d, want %d", num, 3)
	}
}

func Test_newAvlTree(t *testing.T) {
	tree := newAvlTree()
	tree = tree.insert(106465)
	if 106465 != tree.getValByRank(1) {
		t.Errorf("getValByRank() = %d, want %d", tree.getValByRank(1), 106465)
	}
	tree = tree.insert(317721)
	tree = tree.insert(460929)
	tree = tree.insert(644985)
	tree = tree.insert(84185)
	tree = tree.insert(89851)
	var node *avlNode
	node = tree.getNextVal(81968)
	if node.val != 84185 {
		t.Errorf("getNextVal() = %d, want %d", node.val, 84185)
	}
	tree = tree.insert(492737)
	node = tree.getPrevVal(493598)
	if node.val != 492737 {
		t.Errorf("getPrevVal() = %d, want %d", node.val, 492737)
	}
}
