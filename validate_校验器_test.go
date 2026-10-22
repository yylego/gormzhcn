package gormzhcn_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/gormzhcn"
)

func TestT校验器_Validate(t *testing.T) {
	require.NoError(t, gormzhcn.NewT校验器(&T测试产品2{}, gormzhcn.NewT配置项()).Validate())
}

func TestT校验器_Validate_报错(t *testing.T) {
	require.Error(t, gormzhcn.NewT校验器(&T测试产品1{}, gormzhcn.NewT配置项()).Validate())
}
