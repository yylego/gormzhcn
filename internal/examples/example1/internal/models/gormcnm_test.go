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
		&T学生{},
		&T班级{},
	}

	require.True(t, t.Run("GenGormMom", func(t *testing.T) {
		srcPath := runpath.PARENT.Join("models.go")
		t.Log(srcPath)

		p表结构1 := gormzhcn.NewT表结构V2[T学生](srcPath)
		p表结构2 := gormzhcn.NewT表结构V2[T班级](srcPath)
		cfg := gormzhcn.NewT批量编码器([]*gormzhcn.T表结构{
			p表结构1,
			p表结构2,
		}, gormzhcn.NewT配置项().With默认列名样式(gormmomname.NewUppercase63pattern()))
		t.Log(cfg)

		result := cfg.Gen替换源码()
		require.False(t, result.Has有变化())
		require.NoError(t, cfg.Validate验证标签())
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
