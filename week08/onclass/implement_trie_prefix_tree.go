package week08onclass

// https://leetcode-cn.com/problems/implement-trie-prefix-tree/
// leetcode 208. 实现 Trie (前缀树)

//Trie trie struct
type Trie struct {
	count   int
	trieMap map[int]*Trie
}

//Constructor new Trie
func Constructor() Trie {
	return Trie{
		trieMap: make(map[int]*Trie, 5),
	}
}

//Insert 插入
func (t *Trie) Insert(word string) {
	t.find(word, true, false)
	// trie := t
	// for i := 0; i < len(word); i++ {
	// 	cur := int(word[i] - 'a')
	// 	if _, ok := trie.trieMap[cur]; !ok {
	// 		trie.trieMap[cur] = &Trie{
	// 			trieMap: make(map[int]*Trie, 5),
	// 		}
	// 	}
	// 	trie = trie.trieMap[cur]
	// }
	// trie.count++
}

//Search 查找
func (t *Trie) Search(word string) bool {
	return t.find(word, false, false)
	// trie := t
	// for i := 0; i < len(word); i++ {
	// 	cur := int(word[i] - 'a')
	// 	if _, ok := trie.trieMap[cur]; !ok {
	// 		return false
	// 	}
	// 	trie = trie.trieMap[cur]
	// }
	// return trie.count > 0
}

//StartsWith 前缀
func (t *Trie) StartsWith(prefix string) bool {
	return t.find(prefix, false, true)
	// trie := t
	// for i := 0; i < len(prefix); i++ {
	// 	cur := int(prefix[i] - 'a')
	// 	if _, ok := trie.trieMap[cur]; !ok {
	// 		return false
	// 	}
	// 	trie = trie.trieMap[cur]
	// }
	// return true
}

func (t *Trie) find(word string, isInsert, isPrefix bool) bool {
	trie := t
	for i := 0; i < len(word); i++ {
		cur := int(word[i] - 'a')
		if _, ok := trie.trieMap[cur]; !ok {
			if isInsert {
				trie.trieMap[cur] = &Trie{
					trieMap: make(map[int]*Trie, 5),
				}
			} else {
				return false
			}
			
		}
		trie = trie.trieMap[cur]
	}
	if isInsert {
		trie.count++
	}
	if isPrefix {
		return true
	}
	
	return trie.count > 0
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */


//Trie2 trie struct
type Trie2 struct {
	count   int
	arr []*Trie2
}

//Constructor2 new Trie
func Constructor2() Trie2 {
	return Trie2{
		arr:make([]*Trie2, 26),
	}
}

//Insert 插入
func (t *Trie2) Insert(word string) {
	t.find(word, true, false)
}

//Search 查找
func (t *Trie2) Search(word string) bool {
	return t.find(word, false, false)
}

//StartsWith 前缀
func (t *Trie2) StartsWith(prefix string) bool {
	return t.find(prefix, false, true)
}

func (t *Trie2) find(word string, isInsert, isPrefix bool) bool {
	trie := t
	for i := 0; i < len(word); i++ {
		cur := int(word[i] - 'a')
		if trie.arr[cur] == nil {
			if isInsert {
				trie.arr[cur] = &Trie2{
					arr: make([]*Trie2, 26),
				}
			} else {
				return false
			}
		}
		trie = trie.arr[cur]
	}
	if isInsert {
		trie.count++
	}
	if isPrefix {
		return true
	}
	
	return trie.count > 0
}


