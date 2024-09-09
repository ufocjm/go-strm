package go_strm

import (
	"log"
	"os"
	"path/filepath"
)

func GetSuffixName(fileName string) string {
	return filepath.Ext(fileName)[1:]
}

func Contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func CreateFile(fileName, content string) {
	// 创建文件
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error creating file:", err)
		return
	}
	defer file.Close() // 确保在函数退出时关闭文件
	// 写入文本
	_, err = file.WriteString(content)
	if err != nil {
		log.Fatalf("Error writing to file:", err)
		return
	}
}
