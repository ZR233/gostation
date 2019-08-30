/*
@Time : 2019-07-25 10:32
@Author : zr
*/
package utils

import "regexp"

func PhoneIsPhoneNumber(str string) bool {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(str)
}
