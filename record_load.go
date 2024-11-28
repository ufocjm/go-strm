package strm

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

// var executeRecord = map[string]string{}
var executeRecord = sync.Map{}
var recordFile *os.File

func (c *Client) Load() {
	// 打开文件
	var err error
	recordFile, err = os.OpenFile("cache.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
		return
	}
	//defer recordFile.Close()

	// 使用 bufio.Scanner 按行读取文件
	scanner := bufio.NewScanner(recordFile)
	for scanner.Scan() {
		line := scanner.Text()

		// 将每一行以空格分割
		// 使用 SplitN 按空格分割为两部分
		parts := strings.Split(line, " ")

		filePath := strings.Join(parts[:len(parts)-2], " ")
		dateTime := strings.Join(parts[len(parts)-2:], " ")
		// 把路径作为 key，把日期时间作为 value
		executeRecord.Store(filePath, dateTime)
	}

	// 检查扫描过程中是否有错误
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件时出错:", err)
	}

	// 打印map内容
	executeRecord.Range(func(key, value interface{}) bool {
		log.Printf("已处理路径: %s, 日期时间: %s\n", key, value)
		return true

	})
}
