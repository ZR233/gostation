/*
@Time : 2019-07-25 10:42
@Author : zr
*/
package utils

import (
	"math/rand"
	"time"
)

func RandomStr(n int) string {
	str := "abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}
