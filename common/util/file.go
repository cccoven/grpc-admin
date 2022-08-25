package util

import (
	"os"
)

// PathExists 文件是否存在
func PathExists(path string) (bool, error) {
	// os.Stat()：如果文件路径存在，则返回一个描述对象，否则返回错误
	_, err := os.Stat(path)
	if err == nil {
		return true, err
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}
