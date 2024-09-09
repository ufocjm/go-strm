package go_strm

import "log"

func (c *Client) RecordSave(line string) {
	go func() {
		// 写入数据到文件末尾
		_, err := recordFile.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("写入文件时出错: %v", err)
		}
	}()
}
