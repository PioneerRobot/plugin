package plugin

import "github.com/PioneerRobot/core"

//// 插件包解析器

// ParseType 解析类型
type ParseType string

// Parser 插件包解析器
type Parser interface {
	// ParseType 获取该插件包解析器的解析类型
	ParseType () ParseType
	// Read 读取指定文件的内容
	// path 相对于插件包内部的根目录的路径
	Read (path string) ([]byte, error)
	// GetRequire 获取该解析器对应的插件的 Require 配置
	GetRequire () (*Require, error)
	// GetProperty 获取该解析器对应的插件的 PluginProperty 配置
	GetProperty () (*PluginProperty, error)
	// GetComponent 获取组件
	// path 相对于插件包内部的根目录的路径
	GetComponent (path string) (core.Component, error)
	// GetComponents 获取组件
	GetComponents () (*core.ComponentPool, error)
	// Close 关闭读取器
	Close () error
}

// ParserFactory 插件包解析器工厂
// 插件包解析器需要负责构建解析器缓存，减少重复构建解析器带来的消耗
type ParserFactory interface {
	// Supported 返回该插件包工厂支持的插件包解析类型
	Supported () []ParseType
	// GetParser 获取插件包解析器
	// path 插件包的绝对路径
	GetParser (path string) (Parser, error)
	// Destroy 销毁插件包解析器工厂
	// 插件包解析器工厂应该维护其内部的「插件解析器」的关闭
	Destroy () error
}