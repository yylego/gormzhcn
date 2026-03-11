package gormzhcn_test

import (
	"testing"

	"github.com/yylego/gormzhcn"
	"github.com/stretchr/testify/require"
	"github.com/yylego/runpath"
)

type T测试产品1 struct {
	ID  uint    `gorm:"primaryKey"`
	P名称 string  `gorm:"column:name;type:varchar(100);not null;uniqueIndex"`
	P价格 float64 `gorm:"column:price;type:decimal(10,2);not null;index"`
	S库存 int     `gorm:"type:int;default:0"`
	D描述 string  `gorm:"column:description;type:text"`
	S状态 string  `gorm:"column:status;type:varchar(20);default:'active';index"`
}

// TableName specifies the table name of T测试产品1 struct
// TableName 指定 T测试产品1 模型的表名
func (*T测试产品1) TableName() string {
	return "product1s"
}

func TestT批量编码器_Validate验证标签(t *testing.T) {
	objects := []interface{}{
		&T测试产品1{},
	}

	v表结构s := gormzhcn.NewT表结构s(runpath.PARENT.Path(), objects)

	v批量编码器 := gormzhcn.NewT批量编码器(v表结构s, gormzhcn.NewT配置项())
	t.Log(v批量编码器)

	require.Error(t, v批量编码器.Validate验证标签())
}

type T测试产品2 struct {
	ID  uint    `gorm:"primaryKey"`
	P名称 string  `gorm:"column:name;type:varchar(100);not null;uniqueIndex"`
	P价格 float64 `gorm:"column:price;type:decimal(10,2);not null;index"`
	S库存 int     `gorm:"column:stock;type:int;default:0;index"`
	D描述 string  `gorm:"column:description;type:text"`
	S状态 string  `gorm:"column:status;type:varchar(20);default:'active';index"`
}

// TableName specifies the table name of T测试产品2 struct
// TableName 指定 T测试产品2 模型的表名
func (*T测试产品2) TableName() string {
	return "product2s"
}

func TestT批量编码器_Validate验证标签_2(t *testing.T) {
	objects := []interface{}{
		&T测试产品2{},
	}

	v表结构s := gormzhcn.NewT表结构s(runpath.PARENT.Path(), objects)

	v批量编码器 := gormzhcn.NewT批量编码器(v表结构s, gormzhcn.NewT配置项())
	t.Log(v批量编码器)

	require.Error(t, v批量编码器.Validate验证标签())
}
