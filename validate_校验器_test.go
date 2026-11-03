package gormzhcn_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/gormzhcn"
)

func TestNewT校验器(t *testing.T) {
	v校验器 := gormzhcn.NewT校验器(&T产品{}, gormzhcn.NewT配置项())
	require.NotNil(t, v校验器)
	require.NoError(t, v校验器.Validate())
}

func TestT校验器_Validate(t *testing.T) {
	require.NoError(t, gormzhcn.NewT校验器(&T产品{}, gormzhcn.NewT配置项()).Validate())
}

func TestT校验器_Validate_报错(t *testing.T) {
	err := gormzhcn.NewT校验器(&T产品_漏写列名{}, gormzhcn.NewT配置项()).Validate()
	require.ErrorContains(t, err, "struct T产品_漏写列名 field S库存 column")
}
