package week08onclass

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				board: [][]byte{[]byte("XXXX"), []byte("XOOX"), []byte("XXOX"), []byte("XOXX")},
			},
			want: [][]byte{[]byte("XXXX"), []byte("XXXX"), []byte("XXXX"), []byte("XOXX")},
		},
		{
			name: "test02",
			args: args{
				board: [][]byte{[]byte("X")},
			},
			want: [][]byte{[]byte("X")},
		},
		{
			name: "test03",
			args: args{
				board: [][]byte{[]byte("OO"), []byte("OO")},
			},
			want: [][]byte{[]byte("OO"), []byte("OO")},
		},
		{
			name: "test04",
			args: args{
				board: [][]byte{[]byte("OOO"), []byte("OOO"), []byte("OOO")},
			},
			want: [][]byte{[]byte("OOO"), []byte("OOO"), []byte("OOO")},
		},
		{
			name: "test05",
			args: args{
				board: [][]byte{[]byte("XXXXX"), []byte("XOOOX"), []byte("XXOOX"), []byte("XXXOX"), []byte("XOXXX")},
			},
			want: [][]byte{[]byte("XXXXX"), []byte("XXXXX"), []byte("XXXXX"), []byte("XXXXX"), []byte("XOXXX")},
		},
		{
			name: "test06",
			args: args{
				board: [][]byte{[]byte("XOXX"), []byte("OXOX"), []byte("XOXO"), []byte("OXOX"), []byte("XOXO"), []byte("OXOX")},
			},
			want: [][]byte{[]byte("XOXX"), []byte("OXXX"), []byte("XXXO"), []byte("OXXX"), []byte("XXXO"), []byte("OXOX")},
		},
		{
			name: "test07",
			args: args{
				board: [][]byte{
					[]byte("XXXXXXXXXXXXXXXXXXXX"),
					[]byte("XXXXXXXXXOOOXXXXXXXX"),
					[]byte("XXXXXOOOXOXOXXXXXXXX"),
					[]byte("XXXXXOXOXOXOOOXXXXXX"),
					[]byte("XXXXXOXOOOXXXXXXXXXX"),
					[]byte("XXXXXOXXXXXXXXXXXXXX"),
				},
			},
			want: [][]byte{
				[]byte("XXXXXXXXXXXXXXXXXXXX"),
				[]byte("XXXXXXXXXOOOXXXXXXXX"),
				[]byte("XXXXXOOOXOXOXXXXXXXX"),
				[]byte("XXXXXOXOXOXOOOXXXXXX"),
				[]byte("XXXXXOXOOOXXXXXXXXXX"),
				[]byte("XXXXXOXXXXXXXXXXXXXX"),
			},
		},
		{
			name: "test08",
			args: args{
				board: [][]byte{
					[]byte("XOXOXOOOXO"),
					[]byte("XOOXXXOOOX"),
					[]byte("OOOOOOOOXX"),
					[]byte("OOOOOOXOOX"),
					[]byte("OOXXOXXOOO"),
					[]byte("XOOXXXOXXO"),
					[]byte("XOXOOXXOXO"),
					[]byte("XXOXXOXOOX"),
					[]byte("OOOOXOXOXO"),
					[]byte("XXOXXXXOOO"),
				},
			},
			want: [][]byte{
				[]byte("XOXOXOOOXO"),
				[]byte("XOOXXXOOOX"),
				[]byte("OOOOOOOOXX"),
				[]byte("OOOOOOXOOX"),
				[]byte("OOXXOXXOOO"),
				[]byte("XOOXXXXXXO"),
				[]byte("XOXXXXXOXO"),
				[]byte("XXOXXXXOOX"),
				[]byte("OOOOXXXOXO"),
				[]byte("XXOXXXXOOO"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.board); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
