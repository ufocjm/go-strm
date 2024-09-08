package go_strm

import (
	"errors"
	go_alist "github.com/ufocjm/go-alist"
	"log"
	"os"
	"path"
	"path/filepath"
)

func (c *Client) StrmDir() error {
	return c.StrmDirPath(c.config.ScanPath)
}

func (c *Client) StrmDirPath(alistPath string) error {
	if len(alistPath) == 0 {
		alistPath = "/"
	}
	list, err := c.alistClient.List(go_alist.ListReq{
		Path:     alistPath,
		Password: "",
		Page:     1,
		PerPage:  0,
		Refresh:  false,
	})
	if err != nil {
		return err
	}
	if list.Code != 200 {
		return errors.New("alist request error")
	}
	for _, item := range list.Data.Content {
		if item.IsDir {
			// 创建文件夹 然后递归
			localPath := filepath.Join(c.config.OutputPath, alistPath, item.Name)
			// 检查目录是否存在
			if _, err := os.Stat(localPath); os.IsNotExist(err) {
				// 目录不存在，不创建
				log.Printf("创建文件夹 %v", localPath)
				err := os.MkdirAll(localPath, os.ModePerm)
				if err != nil {
					log.Fatalf("创建文件夹失败 %v", err)
					return err
				}
			} else {
				// 目录存在
				log.Printf("目录存在无需创建 %v", filepath.Join(c.config.OutputPath, alistPath, item.Name))
			}
			c.StrmDirPath(path.Join(alistPath, item.Name))
		} else {
			c.Strm(&StrmInfo{
				sign:      item.Sign,
				alistPath: alistPath,
				fileName:  item.Name,
			})
		}
	}
	return nil
}
