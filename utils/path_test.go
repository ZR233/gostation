package utils

import "testing"

func TestGetSrcRoot(t *testing.T) {
	srcRoot := GetSrcRoot()
	print(srcRoot)
}
