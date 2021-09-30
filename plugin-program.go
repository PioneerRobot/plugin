package plugin

// ProgramType 程序类型
type ProgramType string

const (
	// Library 类库
	Library ProgramType = "lib"
)

// Program 程序
type Program struct {
	// 插件类型
	Type ProgramType `json:"type" yaml:"type" xml:"type"`
	// 程序语言
	Language string `json:"language" yaml:"language" xml:"language"`
	// 程序路径
	// 相对于程序包的路径
	Path string `json:"path" yaml:"path" xml:"path"`
	// 加载名称
	// 被加载的变量在类库中的全局变量名
	LoadName string `json:"load-name" yaml:"load-name" xml:"load-name"`
}