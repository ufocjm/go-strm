package go_strm

import (
	"log"
	"os"
)

func (c *Client) CreateFile(fileName, content string) {
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
