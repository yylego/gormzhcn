package gormzhcn

import (
	"github.com/yylego/gormmom"
	"github.com/yylego/gormmom/gormidxname"
	"github.com/yylego/gormmom/gormmomname"
)

// T配置项 represents configuration settings with GORM tag generation via Chinese interface
// Wraps gormmom.Options to provide intuitive Chinese method names
// Contains naming strategies, action controls, and processing options
// Supports customizable tag names, column patterns, and index generation settings
//
// T配置项 代表具有中文接口的 GORM 标签生成配置设置
// 封装 gormmom.Options 以提供直观的中文方法名
// 包含命名策略、行为控制和处理选项
// 支持可自定义的标签名称、列模式和索引生成设置
type T配置项 struct {
	options *gormmom.Options // Native options handler // 内部选项处理器
}

// NewT配置项 creates a new configuration instance with optimized default settings
// Initializes with smart defaults including native language processing
// Returns configured options prepared with Chinese GORM tag generation
//
// NewT配置项 创建新的配置实例，使用优化的默认设置
// 使用包括原生语言处理在内的智能默认值初始化
// 返回为中文 GORM 标签生成准备的配置选项
func NewT配置项() *T配置项 {
	return &T配置项{
		options: gormmom.NewOptions(),
	}
}

// With命名总标签名 sets the system tag name in GORM field processing
// Configures the main tag ID used in struct field tags
// Default is "mom" (native tongue) to support native language
//
// With命名总标签名 设置 GORM 字段处理的系统标签名称
// 配置结构体字段标签中使用的主要标签标识符
// 默认为 "mom"（母语），用于原生语言支持
func (opt *T配置项) With命名总标签名(systemTagName string) *T配置项 {
	opt.options.WithTagName(systemTagName)
	return opt
}

// With列名模式标签 sets the sub-tag name in column naming pattern identification
// Configures the ID used to set column naming strategies
// Default is "mcp" (native tongue column pattern) Unicode field naming
//
// With列名模式标签 设置列命名模式识别的子标签名称
// 配置用于指定列命名策略的标识符
// 默认为 "mcp"（母语列模式），用于 Unicode 字段命名
func (opt *T配置项) With列名模式标签(columnNamingSubTagName string) *T配置项 {
	opt.options.WithSubTagName(columnNamingSubTagName)
	return opt
}

// With列名命名样式 sets the column naming pattern in field processing
// Configures how native language field names are converted to database columns
// Supports various patterns like uppercase, lowercase, and mixed case strategies
//
// With列名命名样式 设置字段处理的列命名模式
// 配置如何将原生语言字段名转换为数据库列名
// 支持各种模式，如大写、小写和混合大小写策略
func (opt *T配置项) With列名命名样式(pattern gormmomname.Pattern) *T配置项 {
	opt.options.WithColumnPattern(pattern)
	return opt
}

// With默认列名样式 sets the default column naming pattern as fallback
// Provides fallback naming plan when specific patterns are not defined
// Ensures consistent column naming across different field types
//
// With默认列名样式 设置作为后备的默认列命名模式
// 在未定义特定模式时提供后备命名策略
// 确保不同字段类型之间的一致列命名
func (opt *T配置项) With默认列名样式(pattern gormmomname.Pattern) *T配置项 {
	opt.options.WithDefaultColumnPattern(pattern)
	return opt
}

// With跳过基本列名 controls if we skip processing simple fields
// When enabled, skips fields without custom naming configurations
// Improves performance through focusing on complex Unicode field names
//
// With跳过基本列名 控制是否跳过处理简单字段
// 启用时，跳过没有自定义命名配置的字段
// 通过专注于复杂的 Unicode 字段名提高性能
func (opt *T配置项) With跳过基本列名(skipBasicColumnName bool) *T配置项 {
	opt.options.WithSkipBasicColumnName(skipBasicColumnName)
	return opt
}

// With重新命名索引 enables disables automatic index name generation
// When enabled, auto-generates database-compatible index names
// Handles native language field names in index naming conventions
//
// With重新命名索引 启用或禁用自动索引名称生成
// 启用时，自动生成数据库兼容的索引名称
// 在索引命名约定中处理原生语言字段名
func (opt *T配置项) With重新命名索引(renewIndexName bool) *T配置项 {
	opt.options.WithAutoIndexName(renewIndexName)
	return opt
}

// With索引命名样式 sets the index naming pattern in database indexes
// Configures how native language field names are converted to index names
// Supports database-optimized naming strategies with improved performance
//
// With索引命名样式 设置数据库索引的索引命名模式
// 配置如何将原生语言字段名转换为索引名称
// 支持为性能优化的数据库命名策略
func (opt *T配置项) With索引命名样式(pattern gormidxname.Pattern) *T配置项 {
	opt.options.WithIndexPattern(pattern)
	return opt
}

// With默认索引样式 sets the default index naming pattern as fallback
// Provides fallback index naming plan consistent database schema
// Ensures each index follows standardized naming conventions
//
// With默认索引样式 设置作为后备的默认索引命名模式
// 为一致的数据库架构提供后备索引命名策略
// 确保所有索引遵循标准化的命名约定
func (opt *T配置项) With默认索引样式(pattern gormidxname.Pattern) *T配置项 {
	opt.options.WithDefaultIndexPattern(pattern)
	return opt
}

// GetOptions returns the underlying gormmom.Options instance
// Provides access to the wrapped options in advanced configuration
// Enables connection with gormmom-based systems
//
// GetOptions 返回底层的 gormmom.Options 实例
// 提供对封装选项的访问以进行高级配置
// 支持与基于 gormmom 的系统的互操作性
func (opt *T配置项) GetOptions() *gormmom.Options {
	return opt.options
}
