package week08onclass

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Constructor();
	trie.Insert("apple");
	if !trie.Search("apple") {
		t.Errorf("can not find apple")
	}
	if trie.Search("app") {
		t.Errorf("find error word app")
	}
	if !trie.StartsWith("app") {
		t.Errorf("can not find prefix app")
	}
	trie.Insert("app")
	if !trie.Search("app") {
		t.Errorf("can not find app")
	}
}

func TestTrie2(t *testing.T) {
	trie := Constructor2();
	trie.Insert("apple");
	if !trie.Search("apple") {
		t.Errorf("can not find apple")
	}
	if trie.Search("app") {
		t.Errorf("find error word app")
	}
	if !trie.StartsWith("app") {
		t.Errorf("can not find prefix app")
	}
	trie.Insert("app")
	if !trie.Search("app") {
		t.Errorf("can not find app")
	}
}
