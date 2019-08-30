/*
@Time : 2019-07-10 10:20
@Author : zr
*/
package utils

import (
	"testing"
)

func TestCountPage(t *testing.T) {
	type args struct {
		totalCount   int64
		onePageCount int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"整除页数", args{3, 1}, 3},
		{"多1页", args{3, 2}, 2},
		{"多10页", args{50, 20}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountPage(tt.args.totalCount, tt.args.onePageCount); got != tt.want {
				t.Errorf("CountPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPageIndex2RowOffset(t *testing.T) {
	type args struct {
		pageIndex    int
		onePageCount int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"第一页", args{1, 10}, 0},
		{"第二页", args{2, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PageIndex2RowOffset(tt.args.pageIndex, tt.args.onePageCount); got != tt.want {
				t.Errorf("PageIndex2RowOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}
