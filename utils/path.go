package utils

import (
	"os"
	"path"
	"runtime"
)

func GetSrcRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	//_, filename, _, _ = runtime.Caller(1)

	pathString := path.Dir(path.Dir(path.Dir(filename)))

	return pathString
}

func MakePath(dirPath string) error {
	_, err := os.Stat(dirPath)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = os.Mkdir(dirPath, os.ModePerm)
	}
	return err
}
