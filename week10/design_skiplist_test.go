package week10

import (
	"reflect"
	"testing"
)

func Test_skiplist(t *testing.T) {
	tests := []struct {
		name string
		want skiplist
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructor() = %v, want %v", got, tt.want)
			}
		})
	}
	obj := constructor()
	obj.add(1)
	obj.add(2)
	obj.add(3)
	if obj.search(0) {
		t.Errorf("search() = %v, want %v", true, false)
	}
	obj.add(4)
	if !obj.search(1) {
		t.Errorf("search() = %v, want %v", false, true)
	}

	if obj.erase(0) {
		t.Errorf("erase() = %v, want %v", true, false)
	}
	if !obj.erase(1) {
		t.Errorf("erase() = %v, want %v", false, true)
	}
	if obj.search(1) {
		t.Errorf("search() = %v, want %v", true, false)
	}

}
