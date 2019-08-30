/*
@Time : 2019-07-11 14:44
@Author : zr
*/
package utils

import (
	"camdig/server/errors"
	"encoding/hex"
	"strings"
)

func MacStringFromBytes(macBytes []byte) (macStr string, err error) {
	if len(macBytes) != 6 {
		err = errors.NewParamErr("mac")
		return
	}

	var mac []string

	for _, v := range macBytes {
		mac = append(mac, hex.EncodeToString([]byte{v}))
	}

	macStr = strings.Join(mac, ":")

	return
}
