/*
@Time : 2019-07-30 16:53
@Author : zr
*/
package utils

import "testing"

func TestMacBytes2String(t *testing.T) {
	type args struct {
		macBytes []byte
	}
	tests := []struct {
		name       string
		args       args
		wantMacStr string
		wantErr    bool
	}{
		{"", args{macBytes: []byte{
			0x1c, 0x2c, 0x3b, 0x3c, 0x8c, 0x4c,
		}}, "1c:2c:3b:3c:8c:4c", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMacStr, err := MacStringFromBytes(tt.args.macBytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("MacStringFromBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMacStr != tt.wantMacStr {
				t.Errorf("MacStringFromBytes() gotMacStr = %v, want %v", gotMacStr, tt.wantMacStr)
			}
		})
	}
}
