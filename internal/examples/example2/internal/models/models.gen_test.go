package models

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/gormcngen"
	"github.com/yylego/gormzhcn"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/runpath/runtestpath"
)

//go:generate go test -run ^TestGenerateModels$ -tags=gormcngen_generate
func TestGenerateModels(t *testing.T) {
	objects := []any{
		&Example{},
	}

	// Validate the naming with gormzhcn first, and generate the columns just when it passes
	// 先用 gormzhcn 校验命名，通过了才生成列
	require.True(t, t.Run("ValidateNaming", func(t *testing.T) {
		require.NoError(t, gormzhcn.NewT批量校验器(objects, gormzhcn.NewT配置项()).Validate())
	}))

	require.True(t, t.Run("GenerateColumns", func(t *testing.T) {
		absPath := osmustexist.FILE(runtestpath.SrcPath(t))
		t.Log(absPath)

		options := gormcngen.NewOptions().
			WithRecvName("T").
			WithTableColumns(false)

		cfg := gormcngen.NewConfigs(objects, options, absPath)
		cfg.WithGoBuildDirective(gormcngen.DirectiveSkipGenerate)
		cfg.Gen()
	}))
}
