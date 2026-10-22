package gormzhcn

import (
	"github.com/yylego/gormmom"
)

// T模型校验器 validates a batch of models through a Chinese interface.
// Wraps gormmom.Configs to validate a batch of models as a group.
//
// T模型校验器 通过中文接口校验一批模型
// 封装 gormmom.Configs 把多个模型作为一组一起校验
type T模型校验器 struct {
	configs gormmom.Configs // Native config batch // 内部配置批次
}

// NewT模型校验器 builds the batch from objects plus a config, parsing each through reflection.
//
// NewT模型校验器 用对象集合和配置构造批量校验器，逐个通过反射解析
func NewT模型校验器(objects []any, v配置项 *T配置项) *T模型校验器 {
	return &T模型校验器{
		configs: gormmom.NewConfigs(objects, v配置项.GetOptions()),
	}
}

// Validate checks table, column and index names across the batch, returning the first issue.
//
// Validate 校验整批模型的表名、列名、索引名，返回遇到的第一个问题
func (c *T模型校验器) Validate() error {
	return c.configs.Validate()
}
