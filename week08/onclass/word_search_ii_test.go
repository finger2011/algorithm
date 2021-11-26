package week08onclass

import (
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				board: [][]byte{
					[]byte("oaan"),
					[]byte("etae"),
					[]byte("ihkr"),
					[]byte("iflv"),
				},
				words: []string{"oath", "pea", "eat", "rain"},
			},
			want: []string{"oath", "eat"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.args.board, tt.args.words); !stringsEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func stringsEqual(str1, str2 []string) bool {
	if len(str1) != len(str2) {
		return false
	}
	var strMap1 = make(map[string]int, len(str1))
	var strMap2 = make(map[string]int, len(str2))
	for _, str := range str2 {
		strMap2[str]++
		strMap1[str]--
	}
	for _, str := range str1 {
		strMap1[str]++
		strMap2[str]--
	}
	for _, val := range strMap1 {
		if val != 0 {
			return false
		}
	}
	for _, val := range strMap2 {
		if val != 0 {
			return false
		}
	}
	return true
}
