package models

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/gormcngen"
	"github.com/yylego/gormzhcn"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/runpath/runtestpath"
)

func TestValidateGormMomAndGenCnm(t *testing.T) {
	objects := []any{
		&T学生{},
		&T班级{},
	}

	require.True(t, t.Run("ValidateGormMom", func(t *testing.T) {
		cfg := gormzhcn.NewT模型校验器(objects, gormzhcn.NewT配置项())
		t.Log(cfg)

		require.NoError(t, cfg.Validate())
	}))

	// 使用 require.True(t, t.Run(---)) 限制只有前一步成功才执行后一步的

	require.True(t, t.Run("GenGormCnm", func(t *testing.T) {
		absPath := osmustexist.FILE(runtestpath.SrcPath(t))
		t.Log(absPath)

		options := gormcngen.NewOptions().
			WithColumnClassExportable(true). //中间类型名称的样式为导出的 T学生Columns
			WithUseTagName(true).
			WithColumnsMethodRecvName("T").
			WithColumnsCheckFieldType(true)

		cfg := gormcngen.NewConfigs(objects, options, absPath)
		cfg.Gen()
	}))
}
