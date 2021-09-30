package plugin

import "github.com/PioneerRobot/core"

// VersionConstraints 版本约束
type VersionConstraints struct {
	// 推荐版本
	Suggest core.Version `json:"suggest" yaml:"suggest" xml:"suggest"`
	// 最小版本约束
	Min core.Version `json:"min" yaml:"min" xml:"min"`
	// 最大版本约束
	Max core.Version `json:"max" yaml:"max" xml:"max"`
}

// Dependency 依赖
type Dependency struct {
	// 插件名
	Name string `json:"name" yaml:"name" xml:"name"`
	// 分组名
	Group string `json:"group" yaml:"group" xml:"group"`
	// 版本约束
	Version VersionConstraints `json:"version" yaml:"version" xml:"version"`
	// 仓库地址
	Repository string `json:"repository" yaml:"repository" xml:"repository"`
}