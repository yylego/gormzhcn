package gormzhcn_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/gormzhcn"
	"gorm.io/gorm/schema"
)

func TestNewT配置项(t *testing.T) {
	// The defaults: GORM's default naming strategy, and a name length cap of 63
	// 默认值：GORM 默认的命名策略，名字长度上限 63
	require.NoError(t, gormzhcn.NewT校验器(&T产品{}, gormzhcn.NewT配置项()).Validate())
}

func TestT配置项_With最大名字长度(t *testing.T) {
	// "idx_products_price", the longest name of T产品, has 18 characters: it passes a cap of 18, and fails a cap of 17
	// T产品 里最长的名字 "idx_products_price" 有 18 个字符：上限 18 能通过，上限 17 就不行
	require.NoError(t, gormzhcn.NewT校验器(&T产品{}, gormzhcn.NewT配置项().With最大名字长度(18)).Validate())
	require.Error(t, gormzhcn.NewT校验器(&T产品{}, gormzhcn.NewT配置项().With最大名字长度(17)).Validate())
}

func TestT配置项_With命名策略(t *testing.T) {
	// Pass the strategy the project runs GORM with, here GORM's default one
	// 传入项目里 GORM 用的命名策略，这里就是 GORM 默认的策略
	opt := gormzhcn.NewT配置项().With命名策略(schema.NamingStrategy{})
	require.NoError(t, gormzhcn.NewT校验器(&T产品{}, opt).Validate())
}
