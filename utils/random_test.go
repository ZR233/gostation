/*
@Time : 2019-07-25 10:42
@Author : zr
*/
package utils

import "testing"

func TestRandomStr(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{5}, 5},
		{"", args{5}, 5},
		{"", args{5}, 5},
	}

	last := ""
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomStr(tt.args.n)
			if len(got) != tt.want {
				t.Errorf("RandomStr() = %v, want %v", got, tt.want)
			}
			if last == got {
				t.Errorf("last = %v, got %v", last, got)
			}

		})
	}

}
