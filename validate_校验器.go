package gormzhcn

import (
	"github.com/yylego/gormmom"
)

// T校验器 validates a single object through a Chinese interface.
// Wraps gormmom.Config to validate one object on its own.
//
// T校验器 通过中文接口校验单个模型
// 封装 gormmom.Config 单独校验一个模型
type T校验器 struct {
	config *gormmom.Config // Native config instance // 内部配置处理器
}

// NewT校验器 prepares a single check from an object plus a config, parsing the object through reflection.
//
// NewT校验器 用对象和配置构造校验器，通过反射解析对象
func NewT校验器(object any, v配置项 *T配置项) *T校验器 {
	return &T校验器{
		config: gormmom.NewConfig(object, v配置项.GetOptions()),
	}
}

// Validate checks one object's table name, column names and index names.
//
// Validate 校验单个模型的表名、列名、索引名
func (c *T校验器) Validate() error {
	return c.config.Validate()
}
