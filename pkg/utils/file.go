package utils

import (
	"os"
)

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}

func MkDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
