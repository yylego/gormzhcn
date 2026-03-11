package models

import (
	"testing"

	"github.com/yylego/gormzhcn"
	"github.com/stretchr/testify/require"
	"github.com/yylego/gormcngen"
	"github.com/yylego/gormmom/gormmomname"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/runpath"
	"github.com/yylego/runpath/runtestpath"
)

func TestGenGormMomAndCnm(t *testing.T) {
	objects := []interface{}{
		&Example{},
	}

	require.True(t, t.Run("GenGormMom", func(t *testing.T) {
		srcPath := runpath.PARENT.Join("models.go")
		t.Log(srcPath)

		param := gormzhcn.NewT表结构V2[Example](srcPath)

		cfg := gormzhcn.NewT编码器(param, gormzhcn.NewT配置项().With默认列名样式(gormmomname.NewUppercase63pattern()))
		t.Log(cfg)

		cfg.Gen替换源码()
	}))

	// 使用 require.True(t, t.Run(---)) 限制只有前一步成功才执行后一步的

	require.True(t, t.Run("GenGormCnm", func(t *testing.T) {
		absPath := osmustexist.FILE(runtestpath.SrcPath(t))
		t.Log(absPath)

		options := gormcngen.NewOptions().
			WithColumnClassExportable(true). //中间类型名称的样式为导出的 ExampleColumns
			WithColumnsMethodRecvName("T").
			WithColumnsCheckFieldType(true)

		cfg := gormcngen.NewConfigs(objects, options, absPath)
		cfg.Gen()
	}))
}
