package gormzhcn

import (
	"github.com/yylego/gormmom"
	"gorm.io/gorm/schema"
)

// T配置项 represents validation settings with a Chinese interface.
// Wraps gormmom.Options to provide intuitive Chinese method names.
//
// T配置项 代表带中文接口的校验设置
// 封装 gormmom.Options 以提供直观的中文方法名
type T配置项 struct {
	options *gormmom.Options // Native options instance // 内部选项处理器
}

// NewT配置项 creates a new configuration with default settings.
//
// NewT配置项 创建带默认设置的新配置
func NewT配置项() *T配置项 {
	return &T配置项{
		options: gormmom.NewOptions(),
	}
}

// With最大名字长度 sets the name length cap and returns the config to chain.
// Databases reject a table, column, index name beyond this length.
//
// With最大名字长度 设置名字长度上限，返回自身以便链式调用
// 超过该长度的表名、列名或索引名数据库不接受
func (opt *T配置项) With最大名字长度(maxNameLength int) *T配置项 {
	opt.options.WithMaxNameLength(maxNameLength)
	return opt
}

// With命名策略 sets the GORM naming strategy and returns the config to chain.
// Set it when a project runs GORM with a custom strategy, so validation matches runtime.
//
// With命名策略 设置 GORM 命名策略，返回自身以便链式调用
// 当项目用了自定义命名策略时设置它，让校验结果和运行时一致
func (opt *T配置项) With命名策略(namingStrategy schema.Namer) *T配置项 {
	opt.options.WithNamingStrategy(namingStrategy)
	return opt
}

// GetOptions returns the underlying gormmom.Options instance.
//
// GetOptions 返回底层的 gormmom.Options 实例
func (opt *T配置项) GetOptions() *gormmom.Options {
	return opt.options
}
