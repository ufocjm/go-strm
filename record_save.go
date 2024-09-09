package go_strm

import "log"

func (c *Client) Clear() {
	log.Printf("清理已处理缓存")
	executeRecord.Clear()
}
