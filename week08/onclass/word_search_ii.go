package week08onclass

// https://leetcode-cn.com/problems/word-search-ii/
// leetcode 212. 单词搜索 II
var visitedWord [][]bool
var trie Trie
var dx = []int{1, 0, 0, -1}
var dy = []int{0, 1, -1, 0}
var ans map[string]int
var str []byte
var wordBoard [][]byte

func findWords(board [][]byte, words []string) []string {
	m := len(board)
	n := len(board[0])
	visitedWord = make([][]bool, m)
	for i := 0; i < m; i++ {
		visitedWord[i] = make([]bool, n)
	}
	trie = Constructor()
	for i := 0; i < len(words); i++ {
		trie.Insert(words[i])
	}
	ans = make(map[string]int, len(words))
	str = make([]byte, 0)
	wordBoard = board
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			visitedWord[i][j] = true
			wordDfs(i, j, &trie)
			visitedWord[i][j] = false
		}
	}
	ansStr := []string{}
	for s := range ans {
		ansStr = append(ansStr, s)
	}
	return ansStr
}

func wordDfs(i, j int, trie *Trie) {
	if trie.trieMap == nil || len(trie.trieMap) == 0 {
		return
	}
	ch := wordBoard[i][j]
	next, ok := trie.trieMap[int(ch-'a')]
	if !ok {
		return
	}
	str = append(str, ch)
	if next.count > 0 {
		ans[string(str)]++
	}
	if next.trieMap == nil || len(next.trieMap) == 0 {
		delete(trie.trieMap, int(ch-'a'))
	} else {
		for k := 0; k < 4; k++ {
			nx, ny := i+dx[k], j+dy[k]
			if nx < 0 || ny < 0 || nx >= len(wordBoard) || ny >= len(wordBoard[0]) || visitedWord[nx][ny] {
				continue
			}
			visitedWord[nx][ny] = true
			wordDfs(nx, ny, next)
			visitedWord[nx][ny] = false
		}
	}

	str = str[0 : len(str)-1]
	return
}
