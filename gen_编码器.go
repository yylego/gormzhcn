package gormzhcn

import (
	"github.com/yylego/gormmom"
)

// T编码器 represents a GORM tag generation engine with Chinese interface
// Wraps gormmom.Config to provide intuitive Chinese method names
// Handles single struct tag generation and source code replacement
// Provides controlled access to tag generation workflow
//
// T编码器 代表具有中文接口的 GORM 标签生成编码器
// 封装 gormmom.Config 以提供直观的中文方法名
// 处理单个结构体标签生成和源代码替换
// 提供对标签生成工作流的控制访问
type T编码器 struct {
	config *gormmom.Config // Native config handler // 内部配置处理器
}

// NewT编码器 creates a new engine instance with table structure and configuration
// Combines table structure information with generation options
// Returns configured engine set up to handle GORM tag generation
// Initializes the tag generation workflow with Chinese interface
//
// NewT编码器 使用表结构和配置创建新的编码器实例
// 结合表结构信息和生成选项
// 返回为 GORM 标签生成准备的配置编码器
// 使用中文接口初始化标签生成工作流
func NewT编码器(v表结构 *T表结构, v配置项 *T配置项) *T编码器 {
	return &T编码器{
		config: gormmom.NewConfig(v表结构.gormStruct, v配置项.options),
	}
}

// Gen替换源码 generates new GORM tags and replaces the source file
// Processes struct definition to add database-compatible tags and column names
// Auto-formats generated code and writes back to source when changes detected
// Provides direct source file modification with Chinese method naming
//
// Gen替换源码 生成新的 GORM 标签并替换原始源文件
// 处理结构体定义以添加数据库兼容的标签和列名
// 检测到变化时自动格式化生成的代码并写回源文件
// 通过中文方法命名提供直接源文件修改
func (cfg *T编码器) Gen替换源码() {
	cfg.config.Generate()
}

// Get预览代码 generates new code with GORM tags without modifying the source file
// Returns generated code result including new source and change information
// Allows preview of changes before file modification takes place
// Provides safe code generation with inspection and validation
//
// Get预览代码 生成带有 GORM 标签的新代码而不修改原始文件
// 返回生成的代码结果，包括新源代码和变化信息
// 允许在实际文件修改前预览变化
// 为检查和验证提供安全的代码生成
func (cfg *T编码器) Get预览代码() *R替换的结果 {
	return newR替换的结果(cfg.config.Preview())
}

// GetConfig returns the underlying gormmom.Config instance
// Provides access to the wrapped config in advanced operations
// Enables connection with gormmom-based systems
// Allows direct access to native configuration when needed
//
// GetConfig 返回底层的 gormmom.Config 实例
// 提供对封装配置的访问以进行高级操作
// 支持与基于 gormmom 的系统的互操作性
// 在需要时允许直接访问内部配置
func (cfg *T编码器) GetConfig() *gormmom.Config {
	return cfg.config
}
