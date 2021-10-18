package week04

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	var twitter = Constructor()
	twitter.PostTweet(1, 5)
	if got := twitter.GetNewsFeed(1); !reflect.DeepEqual(got, []int{5}) {
		t.Errorf("Constructor() = %v, want %v", got, []int{5})
	}
	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	if got := twitter.GetNewsFeed(1); !reflect.DeepEqual(got, []int{6, 5}) {
		t.Errorf("Constructor() = %v, want %v", got, []int{6, 5})
	}
	twitter.Unfollow(1, 2)
	if got := twitter.GetNewsFeed(1); !reflect.DeepEqual(got, []int{5}) {
		t.Errorf("Constructor() = %v, want %v", got, []int{5})
	}
	twitter.PostTweet(1, 7)
	if got := twitter.GetNewsFeed(1); !reflect.DeepEqual(got, []int{7, 5}) {
		t.Errorf("Constructor() = %v, want %v", got, []int{7, 5})
	}
	twitter.PostTweet(1, 4)
	if got := twitter.GetNewsFeed(1); !reflect.DeepEqual(got, []int{4, 7, 5}) {
		t.Errorf("Constructor() = %v, want %v", got, []int{4, 7, 5})
	}
}
