package gormzhcn_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/gormzhcn"
)

// T测试产品1 的母语字段 S库存 没给 column，列名会落到非 ASCII。
type T测试产品1 struct {
	ID  uint
	S库存 int // 无 column → 列名非 ASCII
}

func (*T测试产品1) TableName() string { return "product1s" }

// T测试产品2 的列名、索引名、表名都显式给了数据库安全名字。
type T测试产品2 struct {
	ID  uint    `gorm:"primaryKey"`
	P名称 string  `gorm:"column:name;uniqueIndex:udx_product2_name"`
	P价格 float64 `gorm:"column:price;index:idx_product2_price"`
}

func (*T测试产品2) TableName() string { return "product2s" }

func TestT模型校验器_Validate(t *testing.T) {
	err := gormzhcn.NewT模型校验器([]interface{}{&T测试产品2{}}, gormzhcn.NewT配置项()).Validate()
	require.NoError(t, err)
}

func TestT模型校验器_Validate_报错(t *testing.T) {
	err := gormzhcn.NewT模型校验器([]interface{}{&T测试产品1{}}, gormzhcn.NewT配置项()).Validate()
	require.Error(t, err)
}

func TestT校验器_Validate(t *testing.T) {
	require.NoError(t, gormzhcn.NewT校验器(&T测试产品2{}, gormzhcn.NewT配置项()).Validate())
}
