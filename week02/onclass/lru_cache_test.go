package week02onclass

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // 缓存是 {1=1}
	
	lRUCache.Put(2, 2) // 缓存是 {1=1, 2=2}
	
	a := lRUCache.Get(1) // 返回 1
	if a != 1 {
		t.Errorf("lru cache => get error = get %v, want %v", a, 1)
	}
	
	lRUCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	
	a = lRUCache.Get(2) // 返回 -1 (未找到)
	if a != -1 {
		t.Errorf("lru cache => get error = get %v, want %v", a, -1)
	}
	
	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	
	a = lRUCache.Get(1) // 返回 -1 (未找到)
	if a != -1 {
		t.Errorf("lru cache => get error = get %v, want %v", a, -1)
	}
	a = lRUCache.Get(3) // 返回 3
	if a != 3 {
		t.Errorf("lru cache => get error = get %v, want %v", a, 3)
	}
	
	a = lRUCache.Get(4) // 返回 4
	if a != 4 {
		t.Errorf("lru cache => get error = get %v, want %v", a, 4)
	}
}
