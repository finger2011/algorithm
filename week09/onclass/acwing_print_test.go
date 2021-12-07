package week09onclass

import (
	"reflect"
	"testing"
)

func Test_eightPrint(t *testing.T) {
	type args struct {
		start string
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
				start: "123x46758",
			},
			want: []string{"12345678x", "1234567x8", "1234x6758"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eightPrint(tt.args.start); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eightPrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
