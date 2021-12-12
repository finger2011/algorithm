package week10

// avl平衡二叉树实现

type avlNode struct {
	val    int
	height int // 树高度
	cnt    int //副本数
	size   int // 左右子树的数量
	left   *avlNode
	right  *avlNode
}

func newAvlTree() *avlNode {
	return nil
}

func newAvlNode(val int) *avlNode {
	return &avlNode{
		val:    val,
		height: 1,
		cnt:    1,
		size:   1,
	}
}

func (a *avlNode) getRankByVal(val int) int {
	if a == nil {
		return 0
	}
	var left int
	if a.left != nil {
		left = a.left.size
	}
	if val == a.val {
		return left + 1
	}
	if val < a.val {
		return a.left.getRankByVal(val)
	}
	return a.right.getRankByVal(val) + left + 1
}

func (a *avlNode) getValByRank(rank int) int {
	if a == nil {
		return -1
	}
	left := 0
	if a.left != nil {
		left = a.left.size
	}
	if rank <= left {
		return a.left.getValByRank(rank)
	}
	if rank <= left+a.cnt {
		return a.val
	}
	return a.right.getValByRank(rank - left - a.cnt)
}

func (a *avlNode) find(val int) *avlNode {
	if a == nil {
		return nil
	}
	if a.val == val {
		return a
	}
	if val < a.val {
		return a.left.find(val)
	}
	return a.right.find(val)
}

func (a *avlNode) getPrevVal(val int) *avlNode {
	if a == nil {
		return nil
	}
	if a.val == val {
		return a
	}
	if val < a.val {
		return a.left.getPrevVal(val)
	}
	right := a.right.getPrevVal(val)
	if right != nil && right.val > a.val {
		return right
	}
	return a
}

func (a *avlNode) getNextVal(val int) *avlNode {
	if a == nil {
		return nil
	}
	if a.val == val {
		return a
	}
	if val > a.val {
		return a.right.getNextVal(val)
	}
	left := a.left.getNextVal(val)
	if left != nil && left.val < a.val {
		return left
	}
	return a
}

func (a *avlNode) checkFactor() bool {
	left, right := 0, 0
	if a.left != nil {
		if !a.left.checkFactor() {
			// fmt.Printf("error:%d\n", a.left.val)
			return false
		}
		left = a.left.height
	}
	if a.right != nil {
		if !a.right.checkFactor() {
			// fmt.Printf("error:%d\n", a.right.val)
			return false
		}
		right = a.right.height
	}
	a.update()
	return left-right > -2 && left-right < 2
}

func (a *avlNode) getMax() int {
	if a.right != nil {
		return a.right.getMax()
	}
	return a.val
}

func (a *avlNode) delete(val int) *avlNode {
	if a == nil {
		return a
	}
	if a.val == val {
		if a.cnt > 1 {
			a.cnt--
		} else if a.left != nil && a.right != nil {
			pre := a.left.findPre(a)
			a.val = pre.val
			a.cnt = pre.cnt
			pre.val = val
			pre.cnt = 0
			a.left = a.left.delete(val)
		} else if a.left != nil {
			return a.left
		} else if a.right != nil {
			return a.right
		} else {
			return nil
		}
	} else if val < a.val {
		a.left = a.left.delete(val)
	} else {
		a.right = a.right.delete(val)
	}
	a.update()
	a = a.changeFactor()
	return a
}

func (a *avlNode) insert(val int) *avlNode {
	if a == nil {
		a = newAvlNode(val)
		return a
	}
	if a.val == val {
		a.cnt++
		return a
	}
	if a.val > val {
		a.left = a.left.insert(val)
	} else {
		a.right = a.right.insert(val)
	}
	a.update()
	a = a.changeFactor()
	return a
}

func (a *avlNode) findPre(node *avlNode) *avlNode {
	var ans = a
	if a.right != nil {
		right := a.right.findPre(node)
		if right.val > ans.val {
			ans = right
		}
	}
	if a.left != nil {
		left := a.left.findPre(node)
		if left.val > ans.val {
			ans = left
		}
	}
	return ans
}

func (a *avlNode) changeFactor() *avlNode {
	switch a.getFactor() {
	case 2:
		switch a.left.getFactor() {
		case 1:
			// LL: 一次右旋
			a = a.zig()
		case -1:
			// LR: 先左旋，再右旋
			a.left = a.left.zag()
			a = a.zig()
		case 0:
			// 删除后左子树factor = 0, height = 3, 右子树factor = 0, height = 1
			// 此时 左子树的左右子树高均为2
			//
			q := a.left
			a.left = q.right
			q.right = a
			a.update()
			q.update()
			a = q
		}
	case -2:
		switch a.right.getFactor() {
		case 1:
			// RL: 先右旋，再左旋
			a.right = a.right.zig()
			a = a.zag()
		case -1:
			// RR 一次左旋
			a = a.zag()
		case 0:
			q := a.right
			a.right = q.left
			q.left = a
			a.update()
			q.update()
			a = q
		}

	default:
	}
	return a
}

func (a *avlNode) getFactor() int {
	left, right := 0, 0
	if a.left != nil {
		left = a.left.height
	}
	if a.right != nil {
		right = a.right.height
	}
	return left - right
}

func (a *avlNode) update() {
	left, right := 0, 0
	l, r := 0, 0
	if a.left != nil {
		left = a.left.height
		l = a.left.size
	}
	if a.right != nil {
		right = a.right.height
		r = a.right.size
	}
	a.height = max(left, right) + 1
	a.size = l + r + a.cnt
}

// a的右孩子绕a向左旋转
func (a *avlNode) zag() *avlNode {
	q := a.right
	a.right = q.left
	q.left = a
	a.update()
	q.update()
	return q
}

// a的左孩子绕a向右旋转
func (a *avlNode) zig() *avlNode {
	q := a.left
	a.left = q.right
	q.right = a
	a.update()
	q.update()
	return q
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
