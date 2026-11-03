// Package example1 shows gormzhcn with Chinese fields and cnm tag aliases: the structs pass the check, then work with GORM through the generated columns
//
// example1 演示中文字段加 cnm 标签别名的 gormzhcn 用法：结构体通过校验，之后用生成的列照常配合 GORM 使用
package example1

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yylego/done"
	"github.com/yylego/gormrepo"
	"github.com/yylego/gormrepo/gormclass"
	"github.com/yylego/gormzhcn/internal/examples/example1/internal/models"
	"github.com/yylego/neatjson/neatjsons"
	"github.com/yylego/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var repo学生 = gormrepo.NewRepo(gormclass.Use(&models.T学生{}))
var repo班级 = gormrepo.NewRepo(gormclass.Use(&models.T班级{}))

var caseDB *gorm.DB

func TestMain(m *testing.M) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	done.Done(db.AutoMigrate(
		&models.T学生{},
		&models.T班级{},
	))
	done.Done(db.Save(&models.T班级{
		V班级编码: "LOVE",
		V班级名称: "情侣恋爱班级",
		V班主任名: "月老",
	}).Error)
	done.Done(db.Save(&models.T班级{
		V班级编码: "HANDSOME",
		V班级名称: "型男明星班级",
		V班主任名: "玉帝",
	}).Error)

	done.Done(db.Save(&models.T学生{
		ClassCode: "LOVE",
		Name:      "杨亦乐",
		V性别:       models.M男性,
		V生日:       "1990-08-08",
	}).Error)
	done.Done(db.Save(&models.T学生{
		ClassCode: "LOVE",
		Name:      "李小美",
		V性别:       models.F女性,
		V生日:       "1992-05-20",
	}).Error)
	done.Done(db.Save(&models.T学生{
		ClassCode: "HANDSOME",
		Name:      "张小帅",
		V性别:       models.M男性,
		V生日:       "1991-10-21",
	}).Error)

	caseDB = db
	m.Run()
}

func TestFirst_按名字(t *testing.T) {
	res, err := repo学生.With(t.Context(), caseDB).First(func(db *gorm.DB, cls *models.T学生Columns) *gorm.DB {
		return db.Where(cls.V名字.Eq("杨亦乐"))
	})
	require.NoError(t, err)
	require.Equal(t, "LOVE", res.ClassCode)
	t.Log(neatjsons.S(res))
}

func TestFirst_按名字和性别(t *testing.T) {
	res, err := repo学生.With(t.Context(), caseDB).First(func(db *gorm.DB, cls *models.T学生Columns) *gorm.DB {
		return db.Where(cls.V名字.Eq("李小美")).Where(cls.V性别.Eq(models.F女性))
	})
	require.NoError(t, err)
	require.Equal(t, "1992-05-20", res.V生日)
	t.Log(neatjsons.S(res))
}

func TestFirst_查不到(t *testing.T) {
	res, err := repo学生.With(t.Context(), caseDB).First(func(db *gorm.DB, cls *models.T学生Columns) *gorm.DB {
		return db.Where(cls.V名字.Eq("张小帅")).Where(cls.V性别.Eq(models.F女性))
	})
	require.ErrorIs(t, err, gorm.ErrRecordNotFound)
	require.Nil(t, res)
}

func TestFind_按班级(t *testing.T) {
	v恋爱班级, err := repo班级.With(t.Context(), caseDB).First(func(db *gorm.DB, cls *models.T班级Columns) *gorm.DB {
		return db.Where(cls.V班级编码.Eq("LOVE"))
	})
	require.NoError(t, err)
	require.Equal(t, "月老", v恋爱班级.V班主任名)

	students, err := repo学生.With(t.Context(), caseDB).Find(func(db *gorm.DB, cls *models.T学生Columns) *gorm.DB {
		return db.Where(cls.V班级编码.Eq(v恋爱班级.V班级编码))
	})
	require.NoError(t, err)
	require.Len(t, students, 2)
	t.Log(neatjsons.S(students))
}
