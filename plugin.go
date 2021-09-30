package plugin

import "github.com/PioneerRobot/core"

// PluginProperty 插件属性
//goland:noinspection GoNameStartsWithPackageName
type PluginProperty struct {
	// 插件名
	Name string `json:"name" yaml:"name" xml:"name"`
	// 分组名
	Group string `json:"group" yaml:"group" xml:"group"`
	// 插件描述
	Desc map[string]string `json:"desc" yaml:"desc" xml:"desc"`
	// 版本
	Version core.Version `json:"version" yaml:"version" xml:"version"`
	// 显示名
	DisplayName map[string]string `json:"display-name" yaml:"display-name" xml:"display-name"`
	// 仓库地址
	Repository string `json:"repository" yaml:"repository" xml:"repository"`
	// 程序
	Programs []Program `json:"programs" yaml:"programs" xml:"programs"`
	// 资源文件
	Resources []Resource `json:"resources" yaml:"resources" xml:"resources"`
	// 依赖
	Dependencies []Dependency `json:"dependencies" yaml:"dependencies" xml:"dependencies"`
}