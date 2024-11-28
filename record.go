package strm

import "time"

func (c *Client) Record(alistPath string) {
	executeRecord.Store(alistPath, time.Now().Format("2006-01-02 15:04:05"))
	value, _ := executeRecord.Load(alistPath)
	c.RecordSave(alistPath + " " + value.(string))
}

func (c *Client) Executed(alistPath string) bool {
	_, ok := executeRecord.Load(alistPath)
	return ok
}
