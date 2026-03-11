// Package gormzhcn: Chinese localization package GORM column naming and tag management
// Provides intuitive Chinese function names and interfaces with GORM schema operations
// Built on top of gormmom to provide native Chinese development experience
// Supports automatic tag generation, validation, and schema processing
//
// gormzhcn: GORM 列名和标签管理的中文本地化封装
// 为 GORM 模式操作提供直观的中文函数名和接口
// 基于 gormmom 构建，提供原生中文开发体验
// 支持自动标签生成、验证和模式处理
package gormzhcn

import (
	"github.com/yylego/gormmom"
	"gorm.io/gorm/schema"
)

// T表结构 represents a GORM table structure with Chinese naming conventions
// Wraps gormmom.GormStruct to provide Chinese-localized interface
// Maintains complete source functions while offering intuitive Chinese method names
//
// T表结构 代表具有中文命名约定的 GORM 表结构
// 封装 gormmom.GormStruct 以提供中文本地化接口
// 保持所有原始功能的同时提供直观的中文方法名
type T表结构 struct {
	gormStruct *gormmom.GormStruct // Native GORM struct handler // 内部 GORM 结构处理器
}

// NewT表结构 creates a new table structure instance with basic parameters
// Initializes T表结构 using source path, struct name, and GORM schema
// Provides direct mapping to gormmom.NewGormStruct maximum connection
//
// NewT表结构 使用基本参数创建新的表结构实例
// 使用源路径、结构体名称和 GORM 模式初始化 T表结构
// 提供对 gormmom.NewGormStruct 的直接映射以获得最大兼容性
func NewT表结构(sourcePath string, structName string, gormSchema *schema.Schema) *T表结构 {
	return &T表结构{
		gormStruct: gormmom.NewGormStruct(sourcePath, structName, gormSchema),
	}
}

// NewT表结构V2 creates table structure using generics with type checks
// Uses generic type T to auto-detect struct information
// Note: T must be the type name itself, not a point type
// Provides compile-time type checks and eliminates hand-written type specification
//
// NewT表结构V2 使用泛型创建表结构以确保类型安全
// 使用泛型类型参数 T 自动推断结构体信息
// 注意：T 必须是类型名称本身，而非指针类型
// 提供编译时类型安全并消除手动类型指定
func NewT表结构V2[T any](sourcePath string) *T表结构 {
	return &T表结构{
		gormStruct: gormmom.ParseStruct[T](sourcePath),
	}
}

// NewT表结构V3 creates table structure from runtime object instance
// Accepts both object values and object points with maximum coverage
// Uses reflection to extract type information from the provided object
// Designed to handle dynamic scenarios where type is determined at runtime
//
// NewT表结构V3 从运行时对象实例创建表结构
// 接受对象值和对象指针以获得最大灵活性
// 使用反射从提供的对象中提取类型信息
// 适用于在运行时确定类型的动态场景
func NewT表结构V3(sourcePath string, object interface{}) *T表结构 {
	return &T表结构{
		gormStruct: gormmom.ParseObject(sourcePath, object),
	}
}

// NewT表结构s creates multiple table structures from a collection of objects
// Batch processes multiple objects to generate corresponding T表结构 instances
// Maintains sequence and provides efficient bulk initialization
// Helps with processing multiple related models in a single operation
//
// NewT表结构s 从对象集合创建多个表结构
// 批量处理多个对象以生成相应的 T表结构 实例
// 保持顺序并提供高效的批量初始化
// 适用于在单个操作中处理多个相关模型
func NewT表结构s(root string, objects []interface{}) []*T表结构 {
	structSchemas := gormmom.ParseObjects(root, objects)

	results := make([]*T表结构, 0, len(structSchemas))
	for _, x := range structSchemas {
		results = append(results, &T表结构{
			gormStruct: x, // Wrap gormmom struct with Chinese interface // 用中文接口封装 gormmom 结构
		})
	}
	return results
}
