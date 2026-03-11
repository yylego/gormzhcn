package gormzhcn

import (
	"github.com/yylego/gormmom"
	"github.com/yylego/must"
)

// T批量编码器 represents a batch GORM tag generation engine with Chinese interface
// Manages multiple T编码器 instances to process multiple structs at the same time
// Provides efficient bulk tag generation and source code modification
// Coordinates batch operations with unified result reporting
//
// T批量编码器 代表具有中文接口的批量 GORM 标签生成编码器
// 管理多个 T编码器 实例以同时处理多个结构体
// 提供高效的批量标签生成和源代码修改
// 通过统一结果报告协调批量操作
type T批量编码器 struct {
	configs []*T编码器 // Collection of separate engines // 个别编码器集合
}

// NewT批量编码器 creates a new batch engine with multiple table structures
// Initializes separate engines each table structure with shared configuration
// Provides efficient batch processing setup with multiple struct tag generation
// Returns configured batch engine set up to perform bulk operations
//
// NewT批量编码器 使用多个表结构创建新的批量编码器
// 使用共享配置为每个表结构初始化个别编码器
// 为多个结构体标签生成提供高效的批量处理设置
// 返回为批量操作准备的配置批量编码器
func NewT批量编码器(v表结构s []*T表结构, v配置项 *T配置项) *T批量编码器 {
	var configs = make([]*T编码器, 0, len(v表结构s))
	for _, v表结构 := range v表结构s {
		configs = append(configs, NewT编码器(v表结构, v配置项))
	}
	return &T批量编码器{
		configs: configs,
	}
}

// R批量替换的结果 represents the result of batch tag generation operations
// Contains separate file results and aggregate statistics
// Provides comprehensive tracking of changes across multiple files
// Enables detailed reporting of batch processing outcomes
//
// R批量替换的结果 代表批量标签生成操作的结果
// 包含个别文件结果和聚合统计信息
// 提供跨多个文件变化的全面跟踪
// 支持批量处理结果的详细报告
type R批量替换的结果 struct {
	S替换的结果 []*R替换的结果 // Separate file results // 个别文件结果
	N替换标签数 int       // Count of changed tags // 更改标签的总数
	N替换文件数 int       // Count of changed files // 更改文件的总数
}

// Has有变化 checks if some changes were made during batch processing
// Returns true if some tags files were modified
// Provides quick validation of batch operation effectiveness
// Helps with condition logic based on change detection
//
// Has有变化 检查批量处理期间是否有任何更改
// 如果任何标签或文件被修改，则返回 true
// 提供批量操作有效性的快速验证
// 用于基于变化检测的条件逻辑
func (R *R批量替换的结果) Has有变化() bool {
	return R.N替换标签数 > 0 || R.N替换文件数 > 0
}

// R替换的结果 represents the result of specific file tag generation
// Contains the generated source code and change tracking information
// Provides detailed feedback on tag replacement operations
// Enables inspection of changes before and post file modification
//
// R替换的结果 代表个别文件标签生成的结果
// 包含生成的源代码和变化跟踪信息
// 提供标签替换操作的详细反馈
// 支持在文件修改前后检查变化
type R替换的结果 struct {
	NewCode []byte // Generated source code content // 生成的源代码内容
	SrcPath string // Source file path // 源文件路径
	N替换标签数  int    // Count of changed tags // 更改标签的数量
}

// newR替换的结果 creates a Chinese-localized result package from gormmom result
// Converts gormmom.CodeResult to R替换的结果 with consistent interface
// Maintains complete source data while providing Chinese field names
// Enables seamless integration between gormmom and gormzhcn layers
//
// newR替换的结果 从 gormmom 结果创建中文本地化的结果封装
// 将 gormmom.CodeResult 转换为 R替换的结果 以获得一致的接口
// 在提供中文字段名的同时保持所有原始数据
// 支持 gormmom 和 gormzhcn 层之间的无缝集成
func newR替换的结果(result *gormmom.CodeResult) *R替换的结果 {
	return &R替换的结果{
		NewCode: result.OutputCode,
		SrcPath: result.SourcePath,
		N替换标签数:  result.ChangedLineCount,
	}
}

// Has有变化 checks if some tags were changed in this specific file
// Returns true if tag replacement operations modified the source
// Provides quick validation with single file processing results
// Helps with condition logic and change reporting
//
// Has有变化 检查此个别文件中是否有任何标签被更改
// 如果标签替换操作修改了源代码，则返回 true
// 为单个文件处理结果提供快速验证
// 用于条件逻辑和变化报告
func (R *R替换的结果) Has有变化() bool {
	return R.N替换标签数 > 0
}

// Gen替换源码 performs batch GORM tag generation and source file replacement
// Processes each configured struct and generates database-compatible tags
// Auto-formats and writes changes back to source files when detected
// Returns comprehensive batch result with specific file details and aggregate statistics
//
// Gen替换源码 执行批量 GORM 标签生成和源文件替换
// 处理所有配置的结构体并生成数据库兼容的标签
// 检测到变化时自动格式化并将变化写回源文件
// 返回包含个别文件详情和聚合统计的全面批量结果
func (c *T批量编码器) Gen替换源码() *R批量替换的结果 {
	result := c.GetConfigs().Generate()
	must.Full(result)

	s替换的结果 := make([]*R替换的结果, 0, len(result.Items))
	for _, one := range result.Items {
		s替换的结果 = append(s替换的结果, newR替换的结果(one))
	}
	return &R批量替换的结果{
		S替换的结果: s替换的结果,
		N替换标签数: result.ChangedLineCount,
		N替换文件数: result.ChangedFileCount,
	}
}

// Get预览代码 generates new code with GORM tags without modifying source files
// Processes each configured struct and returns generated code results
// Allows preview of batch changes before file modification takes place
// Provides safe batch code generation with inspection and validation
//
// Get预览代码 生成带有 GORM 标签的新代码而不修改源文件
// 处理所有配置的结构体并返回生成的代码结果
// 允许在实际文件修改前预览批量变化
// 为检查和验证提供安全的批量代码生成
func (c *T批量编码器) Get预览代码() *R批量替换的结果 {
	result := c.GetConfigs().Preview()
	must.Full(result)

	s替换的结果 := make([]*R替换的结果, 0, len(result.Items))
	for _, one := range result.Items {
		s替换的结果 = append(s替换的结果, newR替换的结果(one))
	}
	return &R批量替换的结果{
		S替换的结果: s替换的结果,
		N替换标签数: result.ChangedLineCount,
		N替换文件数: result.ChangedFileCount,
	}
}

// GetConfigs returns the underlying gormmom.Configs collection
// Converts Chinese engine collection to gormmom-compatible config slice
// Provides access to native configurations advanced batch operations
// Enables connection with gormmom batch processing systems
//
// GetConfigs 返回底层的 gormmom.Configs 集合
// 将中文编码器集合转换为 gormmom 兼容的配置切片
// 提供对内部配置的访问以进行高级批量操作
// 支持与 gormmom 批量处理系统的互操作性
func (c *T批量编码器) GetConfigs() gormmom.Configs {
	var configs = make([]*gormmom.Config, 0, len(c.configs))
	for _, v编码器 := range c.configs {
		configs = append(configs, v编码器.GetConfig())
	}
	return configs
}
