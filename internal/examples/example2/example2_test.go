// Package example2 shows gormzhcn with Chinese fields and no cnm tags: the generated columns keep the field names
//
// example2 演示不写 cnm 标签的中文字段：生成的列沿用字段名
package example2

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yylego/done"
	"github.com/yylego/gormrepo"
	"github.com/yylego/gormrepo/gormclass"
	"github.com/yylego/gormzhcn/internal/examples/example2/internal/models"
	"github.com/yylego/neatjson/neatjsons"
	"github.com/yylego/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var caseDB *gorm.DB

func TestMain(m *testing.M) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	done.Done(db.AutoMigrate(&models.Example{}))

	caseDB = db
	m.Run()
}

func TestUse(t *testing.T) {
	for idx := range 10 {
		one := &models.Example{
			V名称: fmt.Sprintf("示例-%d", idx),
			V性别: "女",
			V年龄: 18 + idx,
			V身高: 160,
			V体重: 50,
		}
		require.NoError(t, caseDB.Create(one).Error)
	}

	repo := gormrepo.NewRepo(gormclass.Use(&models.Example{}))

	t.Run("select-first", func(t *testing.T) {
		example, err := repo.With(t.Context(), caseDB).First(func(db *gorm.DB, cls *models.ExampleColumns) *gorm.DB {
			return db.Where(cls.V名称.Eq("示例-3"))
		})
		require.NoError(t, err)
		require.Equal(t, 21, example.V年龄)
		t.Log(neatjsons.S(example))
	})
	t.Run("select-where-in", func(t *testing.T) {
		examples, err := repo.With(t.Context(), caseDB).Find(func(db *gorm.DB, cls *models.ExampleColumns) *gorm.DB {
			return db.Where(cls.V名称.In([]string{"示例-1", "示例-5", "示例-9"}))
		})
		require.NoError(t, err)
		require.Len(t, examples, 3)
		t.Log(neatjsons.S(examples))
	})
}
