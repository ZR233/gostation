/*
@Time : 2019-08-21 13:47
@Author : zr
*/
package service

import (
	_ "camdig/server/utils/test"
	"testing"
)

func TestAuth_UserInfo(t *testing.T) {

	s := NewAuthService()

	s.UserInfo("0290e1c6c17403ca789516775cf5574c")
}
