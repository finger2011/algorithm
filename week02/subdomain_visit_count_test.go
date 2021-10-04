package week02

import (
	"testing"
)

func Test_subdomainVisits(t *testing.T) {
	type args struct {
		cpdomains []string
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
				cpdomains: []string{"9001 discuss.leetcode.com"},
			},
			want: []string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"},
		},
		{
			name: "test02",
			args: args{
				cpdomains: []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
			},
			want: []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subdomainVisits(tt.args.cpdomains); !stringsEqual(got, tt.want) {
				t.Errorf("subdomainVisits() = %v, want %v", got, tt.want)
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

func Test_subdomainVisits2(t *testing.T) {
	type args struct {
		cpdomains []string
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
				cpdomains: []string{"9001 discuss.leetcode.com"},
			},
			want: []string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"},
		},
		{
			name: "test02",
			args: args{
				cpdomains: []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
			},
			want: []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subdomainVisits2(tt.args.cpdomains); !stringsEqual(got, tt.want) {
				t.Errorf("subdomainVisits2() = %v, want %v", got, tt.want)
			}
		})
	}
}
