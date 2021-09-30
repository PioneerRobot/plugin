package plugin

import "github.com/PioneerRobot/core"

//// 插件加载器

// Loader 插件加载器
type Loader interface {
	// SetPluginRootPath 设置插件包存放的根路径
	// path 插件包存放的根路径，必须是绝对路径
	SetPluginRootPath (path string)
	// RegisterReaderFactory 注册一个插件包解析器工厂
	RegisterReaderFactory (factory ParserFactory)
	// GetAllComponents 获取所有组件
	GetAllComponents () (*core.ComponentPool, error)
	// Destroy 销毁插件加载器
	// 插件加载器应该维护注册到它内部的「插件包解析器工厂」的销毁
	Destroy () error
}