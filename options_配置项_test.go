package gormzhcn_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/gormzhcn"
	"gorm.io/gorm/schema"
)

func TestT配置项_With最大名字长度(t *testing.T) {
	opt := gormzhcn.NewT配置项().With最大名字长度(30)
	require.Equal(t, 30, opt.GetOptions().MaxNameLength)
}

func TestT配置项_With命名策略(t *testing.T) {
	namer := schema.NamingStrategy{TablePrefix: "t_"}
	opt := gormzhcn.NewT配置项().With命名策略(namer)
	require.Equal(t, namer, opt.GetOptions().NamingStrategy)
}
