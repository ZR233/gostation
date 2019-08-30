/*
@Time : 2019-08-01 11:23
@Author : zr
*/
package utils

import (
	"encoding/base64"
	"path"
	"strings"
)

func PicDataToBase64(data []byte, fileName string) (base64Str string) {
	ext := path.Ext(fileName)
	ext = strings.TrimLeft(ext, ".")

	imageBase64 := base64.StdEncoding.EncodeToString(data)
	base64Str = "data:image/" + ext + ";base64," + imageBase64
	return
}
