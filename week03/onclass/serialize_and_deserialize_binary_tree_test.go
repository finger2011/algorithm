package week03onclass

import (
	"finger2011/algggorithm/internel"
	"reflect"
	"testing"
)

func TestCodec_serialize(t *testing.T) {
	var root = internel.CreateTreeByIntLevel([]int{1, 2, 3, internel.MinInt, internel.MinInt, 4, 5})
	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(root)
	ans := deser.deserialize(data)
	if !reflect.DeepEqual(ans.TreeToInt(), root.TreeToInt()) {
		t.Errorf("deserialize() = %v, want %v", ans.TreeToInt(), root.TreeToInt())
	}
}

func TestCodec_serialize2(t *testing.T) {
	var root = internel.CreateTreeByIntLevel([]int{1, 2, 3, internel.MinInt, internel.MinInt, 4, 5})
	ser := Constructor()
	deser := Constructor()
	data := ser.serialize2(root)
	ans := deser.deserialize2(data)
	if !reflect.DeepEqual(ans.TreeToInt(), root.TreeToInt()) {
		t.Errorf("deserialize() = %v, want %v", ans.TreeToInt(), root.TreeToInt())
	}
}
