package plugin

import "github.com/PioneerRobot/core"

// Require 插件包要求规范
type Require struct {
	// 核心库版本号
	Core core.Version `json:"core" yaml:"core" xml:"core"`
	// 插件规范版本号
	Plugin core.Version `json:"plugin" yaml:"plugin" xml:"plugin"`
}