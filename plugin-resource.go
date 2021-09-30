package plugin

// ResourceType 资源类型
type ResourceType string

const (
	// Binary 二进制文件
	Binary ResourceType = "binary"
	// Image 图片
	Image = "image"
	// Audio 音频
	Audio = "audio"
	// Video 视频
	Video = "video"
)

// Resource 资源
type Resource struct {
	// 资源类型
	Type ResourceType `json:"type" yaml:"type" xml:"type"`
	// 资源路径
	// 相对于程序包的路径
	Path string `json:"path" yaml:"path" xml:"path"`
	// 映射路径
	// 映射到 core.Config 中的配置路径
	Mapping string `json:"mapping" yaml:"mapping" xml:"mapping"`
}