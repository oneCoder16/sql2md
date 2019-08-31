package util

import (
	"io/ioutil"
	"os"
)

//获取文件所有内容
func GetContentByFile(filePath string) (content []byte, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	content, err = ioutil.ReadAll(file)
	return
}
