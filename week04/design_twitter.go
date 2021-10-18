package week04

// https://leetcode-cn.com/problems/design-twitter/
// leetcode 355. 设计推特

//Twitter twitter
type Twitter struct {
	userFollows map[int]map[int]int // 用户关注列表，默认关注自己
	userPosts   map[int]*post       // 用户推文
	userFeed    map[int][]int
	id          int
}

type post struct {
	id     int
	userID int
	postID int
	next   *post
}

type binaryLastPost struct {
	heap []*post
}

func (b *binaryLastPost) getLast() *post {
	return b.heap[1]
}
func (b *binaryLastPost) removeLast() {
	b.heap[1] = b.heap[len(b.heap)-1]
	b.heap = b.heap[0 : len(b.heap)-1]
	b.heapifyDown(1)
}

func (b *binaryLastPost) heapifyDown(i int) {
	var ch = i * 2
	for ch < len(b.heap) {
		if ch+1 < len(b.heap) && b.heap[ch+1].id > b.heap[ch].id {
			ch++
		}
		if b.heap[ch].id > b.heap[i].id {
			var tmp = b.heap[ch]
			b.heap[ch] = b.heap[i]
			b.heap[i] = tmp
			i = ch
			ch = i * 2
		} else {
			break
		}
	}
}

func (b *binaryLastPost) insert(p *post) {
	b.heap = append(b.heap, p)
	b.heapifyUp(len(b.heap) - 1)
}

func (b *binaryLastPost) empty() bool {
	return len(b.heap) == 1
}

func (b *binaryLastPost) heapifyUp(i int) {
	for i > 1 {
		if b.heap[i].id > b.heap[i/2].id {
			var tmp = b.heap[i]
			b.heap[i] = b.heap[i/2]
			b.heap[i/2] = tmp
			i /= 2
		} else {
			break
		}
	}
}

//Constructor return new twitter
func Constructor() Twitter {
	return Twitter{
		userFollows: make(map[int]map[int]int, 10),
		userPosts:   make(map[int]*post, 10),
	}
}

//PostTweet post new tweet
func (t *Twitter) PostTweet(userID int, tweetID int) {
	t.id++
	var newPost = &post{
		userID: userID,
		postID: tweetID,
		id:     t.id,
	}
	if len(t.userFollows[userID]) == 0 {
		t.userFollows[userID] = make(map[int]int, 10)
		t.userFollows[userID][userID] = 1
		t.userPosts[userID] = newPost
	} else {
		if t.userPosts[userID] == nil {
			t.userPosts[userID] = newPost
		} else {
			newPost.next = t.userPosts[userID]
			t.userPosts[userID] = newPost
		}
	}
}

//GetNewsFeed get user new feeds
func (t *Twitter) GetNewsFeed(userID int) []int {
	var b = &binaryLastPost{heap: []*post{&post{}}}
	for followerID := range t.userFollows[userID] {
		if t.userPosts[followerID] != nil {
			b.insert(t.userPosts[followerID])
		}
	}
	var ans []int
	for i := 0; i < 10; i++ {
		if !b.empty() {
			var post = b.getLast()
			b.removeLast()
			if post.next != nil {
				b.insert(post.next)
			}
			ans = append(ans, post.postID)
		} else {
			break
		}
	}
	return ans
}
//Follow followerID follow followeeID
func (t *Twitter) Follow(followerID int, followeeID int) {
	if len(t.userFollows[followerID]) == 0 {
		t.userFollows[followerID] = make(map[int]int, 10)
		t.userFollows[followerID][followerID] = 1
	}
	t.userFollows[followerID][followeeID] = 1
}
//Unfollow followerID unfollow followeeID
func (t *Twitter) Unfollow(followerID int, followeeID int) {
	if _, ok := t.userFollows[followerID][followeeID]; ok {
		delete(t.userFollows[followerID], followeeID)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userID,tweetID);
 * param_2 := obj.GetNewsFeed(userID);
 * obj.Follow(followerID,followeeID);
 * obj.Unfollow(followerID,followeeID);
 */
