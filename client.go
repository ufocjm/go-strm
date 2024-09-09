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
	StrmLinkConfig *StrmLinkConfig // http直链 strm内部是http
	StrmFileConfig *StrmFileConfig // 文件直链接 strm内部是/开头的路径文件
}

type StrmLinkConfig struct {
	UrlPrefix string // 直链前缀
}

type StrmFileConfig struct {
	PathPrefix string // 路径前缀，如果最终只用这个路径拼接成为strm文件的内容和alist的路径不一致 请参考embyExternalUrl里的embyPathMapping进行配置
}

func NewClient(config *Config) *Client {
	if config.AlistConfig == nil {
		log.Fatal("alist配置不能为空")
		return nil
	}
	if strings.HasSuffix(config.AlistConfig.ServerUrl, "/") {
		panic("alist地址不能以[/]结尾")
	}
	if len(config.ScanPath) == 0 {
		panic("扫描路径不能为空，至少为[/]")
	}
	if len(config.ScanPath) > 0 && !strings.HasPrefix(config.ScanPath, "/") {
		panic("扫描路径参数异常，必须以[/]开始")
	}
	if len(config.ScanPath) > 1 && strings.HasSuffix(config.ScanPath, "/") {
		panic("扫描路径参数异常，不能以[/]结尾")
	}
	if len(config.OutputPath) == 0 {
		panic("输出路径不能为空")
	}
	if !strings.HasPrefix(config.OutputPath, "/") {
		panic("输出路径必须以[/]开始")
	}
	if strings.HasSuffix(config.OutputPath, "/") {
		panic("输出路径不能以[/]结尾")
	}
	if len(config.VideoSuffix) == 0 {
		// 默认就这俩类型
		config.VideoSuffix = []string{"mp4", "mkv"}
		log.Printf("默认创建影视strm文件后缀为: %v", config.VideoSuffix)
	}
	if config.UseHttp {
		if config.StrmLinkConfig == nil {
			panic("strm类型为http直链 直链url地址不能为空")
		}
		if strings.HasSuffix(config.StrmLinkConfig.UrlPrefix, "/") {
			panic("strm类型为http直链 直链url地址不能以[/]结尾")
		}
	} else {
		if config.StrmFileConfig != nil {
			if config.StrmFileConfig.PathPrefix == "/" {
				panic("输出路径不能只有一个[/]")
			}
			if !strings.HasPrefix(config.StrmFileConfig.PathPrefix, "/") {
				panic("strm类型为file 文件路径必须以[/]开始")
			}
			if strings.HasSuffix(config.StrmFileConfig.PathPrefix, "/") {
				panic("strm类型为file 文件路径不能以[/]结尾")
			}
		}
	}

	return &Client{
		alistClient: go_alist.NewClient(config.AlistConfig),
		config:      config,
	}
}
