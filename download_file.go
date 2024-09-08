package go_strm

import (
	"io"
	"log"
	"net/http"
	"os"
)

func (c *Client) DownloadFile(url, target string) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("下载文件失败 %v %v", url, err)
		return err
	}
	defer resp.Body.Close()
	targetFile, err := os.Create(target)
	if err != nil {
		log.Fatalf("创建文件失败 %v %v", target, err)
		return err
	}
	defer targetFile.Close()
	_, err = io.Copy(targetFile, resp.Body)
	if err != nil {
		log.Fatalf("写入文件失败 %v", err)
		return err
	}
	return nil
}
