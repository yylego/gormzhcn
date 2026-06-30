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
		V学号:       0,
		Name:      "杨亦乐",
		V性别:       models.M男性,
		V生日:       "1990-08-08",
	}).Error)
	done.Done(db.Save(&models.T学生{
		ClassCode: "LOVE",
		V学号:       0,
		Name:      "刘亦菲",
		V性别:       models.F女性,
		V生日:       "1987-08-25",
	}).Error)
	done.Done(db.Save(&models.T学生{
		ClassCode: "HANDSOME",
		V学号:       0,
		Name:      "古天乐",
		V性别:       models.M男性,
		V生日:       "1970-10-21",
	}).Error)

	caseDB = db
	m.Run()
}

func TestSelect(t *testing.T) {
	res, err := repo学生.Repo(caseDB).First(func(db *gorm.DB, cls *models.T学生Columns) *gorm.DB {
		return db.Where(cls.V名字.Eq("杨亦乐"))
	})
	require.NoError(t, err)
	require.Equal(t, "杨亦乐", res.Name)
	t.Log(neatjsons.S(res))
}

func TestSelect_x2x(t *testing.T) {
	res, err := repo学生.Repo(caseDB).First(func(db *gorm.DB, cls *models.T学生Columns) *gorm.DB {
		return db.Where(cls.V名字.Eq("刘亦菲")).Where(cls.V性别.Eq(models.F女性))
	})
	require.NoError(t, err)
	require.Equal(t, "刘亦菲", res.Name)
	t.Log(neatjsons.S(res))
}

func TestSelect_x3x(t *testing.T) {
	res, err := repo学生.Repo(caseDB).First(func(db *gorm.DB, cls *models.T学生Columns) *gorm.DB {
		return db.Where(cls.V名字.Eq("古天乐")).Where(cls.V性别.Eq(models.F女性))
	})
	require.ErrorIs(t, err, gorm.ErrRecordNotFound)
	require.Nil(t, res)
}

func TestSelect_x4x(t *testing.T) {
	v恋爱班级, err := repo班级.Repo(caseDB).First(func(db *gorm.DB, cls *models.T班级Columns) *gorm.DB {
		return db.Where(cls.V班级编码.Eq("LOVE"))
	})
	require.NoError(t, err)
	t.Log(neatjsons.S(v恋爱班级))

	students, err := repo学生.Repo(caseDB).Find(func(db *gorm.DB, cls *models.T学生Columns) *gorm.DB {
		return db.Where(cls.V班级编码.Eq(v恋爱班级.V班级编码))
	})
	require.NoError(t, err)
	require.Len(t, students, 2)
	t.Log(neatjsons.S(students))
}
