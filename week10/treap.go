package week10

// treap实现
import (
	"errors"
	"math/rand"
)

const maxInt = 1024

type treapNode struct {
	key   int // 关键值
	val   int // 随机权值
	cnt   int // 副本数
	size  int // 子树大小
	left  *treapNode
	right *treapNode
}

func newTreapNode(key int) *treapNode {
	return &treapNode{
		key:  key,
		val:  rand.Intn(maxInt),
		cnt:  1,
		size: 1,
	}
}

func newTreap() *treapNode {
	root := newTreapNode(1e9)
	root.right = newTreapNode(-1e9)
	return root
}

func (t *treapNode) getPrev(n int) int {
	if t == nil {
		return 2e9
	}
	if n <= t.key {
		return t.left.getPrev(n)
	}
	right := t.right.getPrev(n)
	if right > t.key {
		return right
	}
	return t.key
}

func (t *treapNode) getNext(val int) int {
	if t == nil {
		return 2e9
	}
	if val >= t.key {
		return t.right.getNext(val)
	}
	left := t.left.getNext(val)
	if left < t.key {
		return left
	}
	return t.key
}

func (t *treapNode) getValByRank(rank int) (int, error) {
	if t == nil {
		return 0, errors.New("error")
	}
	var left int
	if t.left != nil {
		left = t.left.size
	}
	if left >= rank {
		return t.left.getValByRank(rank)
	}
	if rank <= left+t.cnt {
		return t.key, nil
	}
	return t.right.getValByRank(rank - left - t.cnt)
}

func (t *treapNode) getRankByVal(n int) int {
	if t == nil {
		return 0
	}
	var left int
	if t.left != nil {
		left = t.left.size
	}
	if t.key == n {
		return left + 1
	} else if n < t.key {
		return t.left.getRankByVal(n)
	} else {
		return t.right.getRankByVal(n) + left + 1
	}
}

func (t *treapNode) delete(n int) *treapNode {
	if t == nil {
		return nil
	}
	if t.key == n {
		if t.cnt > 1 {
			t.cnt--
			t.update()
			return t
		}
		if t.left == nil && t.right == nil {
			return nil
		}
		res := t
		if (t.right == nil) || (t.left != nil && t.left.key > t.right.key) {
			res = t.zig()
			res.right = res.right.delete(n)
		} else {
			res = t.zag()
			res.left = res.left.delete(n)
		}
		if res != nil {
			res.update()
		}
		return res
	} else if t.key < n {
		t.right = t.right.delete(n)
	} else {
		t.left = t.left.delete(n)
	}
	if t != nil {
		t.update()
	}
	return t
}

func (t *treapNode) insert(n int) *treapNode {
	ans := t
	if t == nil {
		return newTreapNode(n)
	}
	if t.key == n {
		t.cnt++
	} else if t.key < n {
		t.right = t.right.insert(n)
		// 右孩子更新后无法满足堆性质，左旋
		if t.val < t.right.val {
			ans = ans.zag()
		}
	} else {
		t.left = t.left.insert(n)
		if t.val < t.left.val {
			ans = ans.zig()
		}
	}
	return ans
}

// t的右孩子绕t向左旋转
func (t *treapNode) zag() *treapNode {
	q := t.right
	t.right = q.left
	q.left = t
	t.update()
	q.update()
	return q
}

// t的左孩子绕t向右旋转
func (t *treapNode) zig() *treapNode {
	q := t.left
	t.left = q.right
	q.right = t
	t.update()
	q.update()
	return q
}

func (t *treapNode) update() {
	left, right := 0, 0
	if t.left != nil {
		left = t.left.size
	}
	if t.right != nil {
		right = t.right.size
	}
	t.size = left + right + t.cnt
}
