package go_strm

import (
	"github.com/ufocjm/go-alist"
	"log"
	"strings"
)

type Client struct {
	alistClient *go_alist.Client
	config      *Config
}
type Config struct {
	AlistConfig    *go_alist.Config
	ScanPath       string          // 扫描的路径
	OutputPath     string          // 输出路径
	VideoSuffix    []string        // 要扫描的视频后缀名
	UseHttp        bool            // 是否使用http直链
	StrmLinkConfig *StrmLinkConfig // http直链
}

type StrmLinkConfig struct {
	UrlPrefix string // 直链前缀
}

type StrmFileConfig struct {
}

func NewClient(config *Config) *Client {
	if config.AlistConfig == nil {
		log.Fatal("alist配置不能为空")
		return nil
	}
	config.AlistConfig.ServerUrl = strings.TrimSuffix(config.AlistConfig.ServerUrl, "/")
	if len(config.ScanPath) != 0 {
		config.ScanPath = strings.TrimSuffix(config.ScanPath, "/")
	}
	if len(config.OutputPath) == 0 {
		log.Fatal("输出路径不能为空")
		return nil
	} else {
		config.OutputPath = strings.TrimSuffix(config.OutputPath, "/")
	}
	if len(config.VideoSuffix) == 0 {
		// 默认就这俩类型
		config.VideoSuffix = []string{"mp4", "mkv"}
	}
	if len(config.StrmLinkConfig.UrlPrefix) == 0 {
		config.StrmLinkConfig.UrlPrefix = config.AlistConfig.ServerUrl
	}
	config.StrmLinkConfig.UrlPrefix = strings.TrimSuffix(config.StrmLinkConfig.UrlPrefix, "/")
	return &Client{
		alistClient: go_alist.NewClient(config.AlistConfig),
		config:      config,
	}
}
