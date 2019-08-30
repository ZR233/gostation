/*
@Time : 2019-06-26 9:31
@Author : zr
@File : router
@Software: GoLand
*/
package router

import (
	"camdig/server/controller"
	"reflect"
	"testing"
)

func TestNewDefauteRouter(t *testing.T) {
	tests := []struct {
		name string
		want *Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouter_AutoRegisterController(t *testing.T) {
	r := NewDefaultRouter()
	r.AutoRegisterController("/v1", &controller.Auth{})
}
