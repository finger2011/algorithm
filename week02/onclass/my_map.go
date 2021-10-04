package week02onclass

import (
	"errors"
)

// homework:实现map[string][int]

var errNotFound = errors.New("not found")
var errKey = errors.New("key should be string and not number")
var maxLength = 13
var base = uint64(27) //进制
var divisor = uint64(99991) //取模相关

type myMap struct {
	array [99991]*listNode
}

type listNode struct {
	key  string
	val  int
	next *listNode
}

func newMap() *myMap {
	return &myMap{}
}

func (my myMap) find(key string) (int, error) {
	var node, err = my.findNode(key)
	if err != nil {
		return 0, err
	}
	return node.val, nil
}

func (my *myMap) put(key string, val int) error {
	var node, err = my.findNode(key)
	if err == nil {
		node.val = val
	} else if err == errKey {
		return err
	} else {
		var newNode = &listNode{key: key, val: val}
		for node.next != nil {
			node = node.next
		}
		node.next = newNode
	}
	return nil
}

func (my *myMap) remove(key string) error {
	var offset, err = calMyMapHash(key)
	if err != nil {
		return err
	}
	var node = my.array[offset]
	if node == nil {
		return errNotFound
	}
	var prev = node
	node = node.next
	for node != nil {
		if node.key == key {
			prev.next = node.next
			return nil
		}
		prev = node
		node = node.next
	}
	return errNotFound
}

func (my *myMap) findNode(key string) (*listNode, error) {
	var offset, err = calMyMapHash(key)
	if err != nil {
		return nil, err
	}
	var node = my.array[offset]
	if node == nil {
		node = &listNode{}
		my.array[offset] = node
	}
	var prev = node
	node = node.next
	for node != nil {
		if node.key == key {
			return node, nil
		}
		prev = node
		node = node.next
	}
	return prev, errNotFound
}

// 27进制 int64可选择的最大长度为13位，超过13位有可能溢出
func calMyMapHash(key string) (int, error) {
	if "" == key {
		return 0, errKey
	}
	var ant uint64
	var keys []string
	for i := 0; i < len(key); i += maxLength {
		keys = append(keys, key[i:min(i+maxLength, len(key) - 1)])
	}
	for _, str := range(keys) {
		var chNum uint64
		for i := 0; i < len(str); i++ {
			chNum *= base
			if 64 < str[i] && str[i] < 91 {
				//大写字母
				chNum = uint64(str[i] - 64)
			} else if 96 < str[i] && str[i] < 123 {
				chNum = uint64(str[i] - 96)
			} else {
				return 0, errKey
			}
		}
		ant += chNum%divisor
	}
	ant = ant % divisor
	return int(ant), nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
