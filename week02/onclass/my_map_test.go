package week02onclass

import (
	"testing"
)

var largekey = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

func Test_KeyNotExist(t *testing.T) {
	var myMap = newMap()
	var _, err = myMap.find("a")
	if err == nil || err != errNotFound {
		t.Errorf("myMap => find error: %v, want %v", err, errNotFound)
	}
	// large string
	_, err = myMap.find(largekey)
	if err == nil || err != errNotFound {
		t.Errorf("myMap => find error: %v, want %v", err, errNotFound)
	}
}

func Test_ErrorKey(t *testing.T) {
	var myMap = newMap()
	// error key
	err := myMap.put("", 100)
	if err != errKey {
		t.Errorf("myMap => put error: %v, want %v", err, errKey)
	}

	err = myMap.put("ab12cd", 100)
	if err != errKey {
		t.Errorf("myMap => put error: %v, want %v", err, errKey)
	}
}

func Test_PutNewKey(t *testing.T) {
	var myMap = newMap()

	//put and create test
	err := myMap.put("a", 123)
	if err != nil {
		t.Errorf("myMap => put error: %v", err)
	}
	err = myMap.put(largekey, 100)
	if err != nil {
		t.Errorf("myMap => put error: %v", err)
	}
	val, err := myMap.find("a")
	if err != nil {
		t.Errorf("myMap => find error: %v", err)
	}
	if val != 123 {
		t.Errorf("myMap => find error = get %v, want %v", val, 123)
	}
	val, err = myMap.find(largekey)
	if err != nil {
		t.Errorf("myMap => find error: %v", err)
	}
	if val != 100 {
		t.Errorf("myMap => find error = get %v, want %v", val, 100)
	}
}

func Test_PutForUpdateKey(t *testing.T) {
	var myMap = newMap()
	err := myMap.put("a", 234)
	if err != nil {
		t.Errorf("myMap => put error: %v", err)
	}
	// put and update test
	err = myMap.put("a", 456)
	if err != nil {
		t.Errorf("myMap => put error: %v", err)
	}
	val, err := myMap.find("a")
	if err != nil {
		t.Errorf("myMap => find error: %v", err)
	}
	if val != 456 {
		t.Errorf("myMap => find error = get %v, want %v", val, 456)
	}
	// large key
	err = myMap.put(largekey, 100)
	if err != nil {
		t.Errorf("myMap => put error: %v", err)
	}
	err = myMap.put(largekey, 200)
	if err != nil {
		t.Errorf("myMap => put error: %v", err)
	}
	val, err = myMap.find(largekey)
	if err != nil {
		t.Errorf("myMap => find error: %v", err)
	}
	if val != 200 {
		t.Errorf("myMap => find error = get %v, want %v", val, 200)
	}
}

func Test_Remove(t *testing.T) {
	var myMap = newMap()
	// remove test
	err := myMap.remove("b")
	if err != errNotFound {
		t.Errorf("myMap => remove error = get %v, want %v", err, errNotFound)
	}
	myMap.put("a", 123)
	err = myMap.remove("a")
	if err != nil {
		t.Errorf("myMap => remove error = get %v, want %v", err, nil)
	}
	_, err = myMap.find("a")
	if err != errNotFound {
		t.Errorf("myMap => remove error = get %v, want %v", err, errNotFound)
	}

	err = myMap.remove(largekey + "b")
	if err != errNotFound {
		t.Errorf("myMap => remove error = get %v, want %v", err, errNotFound)
	}
	myMap.put(largekey, 100)
	err = myMap.remove(largekey)
	if err != nil {
		t.Errorf("myMap => remove error = get %v, want %v", err, nil)
	}
	_, err = myMap.find(largekey)
	if err != errNotFound {
		t.Errorf("myMap => remove error = get %v, want %v", err, errNotFound)
	}
}

func Test_Collision(t *testing.T) {
	var myMap = newMap()

	//碰撞
	err := myMap.put("jj", 100)
	if err != nil {
		t.Errorf("myMap => put error = get %v, want %v", err, nil)
	}
	err = myMap.put("ik", 200)
	if err != nil {
		t.Errorf("myMap => put error = get %v, want %v", err, nil)
	}
	val, err := myMap.find("jj")
	if err != nil {
		t.Errorf("myMap => find error = get %v, want %v", err, nil)
	}
	if val != 100 {
		t.Errorf("myMap => find error = get %v, want %v", val, 100)
	}
	val, err = myMap.find("ik")
	if err != nil {
		t.Errorf("myMap => find error = get %v, want %v", err, nil)
	}
	if val != 200 {
		t.Errorf("myMap => find error = get %v, want %v", val, 200)
	}
}
