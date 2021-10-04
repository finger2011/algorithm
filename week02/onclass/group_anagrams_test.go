package week02onclass

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name: "test02",
			args: args{
				strs: []string{""},
			},
			want: [][]string{{""}},
		},
		{
			name: "test03",
			args: args{
				strs: []string{"a"},
			},
			want: [][]string{{"a"}},
		},
		{
			name: "test04",
			args: args{
				strs: []string{"abbbbbbbbbbb","aaaaaaaaaaab"},
			},
			want: [][]string{{"abbbbbbbbbbb"}, {"aaaaaaaaaaab"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !strsEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func strsEqual(str1, str2 [][]string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for _, str := range str1 {
		var isContain bool
		for _, compareStr := range str2 {
			if strEqual(str, compareStr) {
				isContain = true
				break
			}
		}
		if !isContain {
			return false
		}
	}
	return true
}

func strEqual(str1, str2 []string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for _, str := range str1 {
		var isContain bool
		for _, compareStr := range str2 {
			if reflect.DeepEqual(str, compareStr) {
				isContain = true
				break
			}
		}
		if !isContain {
			return false
		}
	}
	return true
}
