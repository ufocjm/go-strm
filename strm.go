package go_strm

import (
	"log"
	"path"
	"path/filepath"
	"strings"
)

func (c *Client) Strm(info *StrmInfo) {
	alistFileUrl := c.config.AlistConfig.ServerUrl + path.Join("/d", info.alistPath, info.fileName) + "?sign=" + info.sign
	suffixName := GetSuffixName(info.fileName)
	if Contains(c.config.VideoSuffix, suffixName) {
		log.Printf("处理文件 %v", info.fileName)
		if c.config.UseHttp {
			c.StrmFile(info, alistFileUrl)
		} else {
			strmContent := filepath.Join(c.config.StrmFileConfig.PathPrefix, info.alistPath, info.fileName)
			c.StrmFile(info, strmContent)
		}
	} else {
		// 进行复制文件
		log.Printf("复制文件 %v", info.fileName)
		_ = c.DownloadFile(alistFileUrl, filepath.Join(c.config.OutputPath, info.alistPath, info.fileName))
	}
}

func (c *Client) StrmFile(info *StrmInfo, strmContent string) {
	// 使用 filepath.Ext 获取文件的扩展名
	ext := GetSuffixName(info.fileName)
	// 去掉旧的扩展名并加上新的 .strm 扩展名
	strmFileName := strings.TrimSuffix(info.fileName, ext) + "strm"
	localFilePath := filepath.Join(c.config.OutputPath, info.alistPath, strmFileName)
	CreateFile(localFilePath, strmContent)
}

type StrmInfo struct {
	sign      string // alist 签名
	alistPath string
	fileName  string
}
