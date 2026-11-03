package gormzhcn_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/gormzhcn"
)

// T产品 的列名、索引名、表名都显式给了数据库安全名字。
type T产品 struct {
	ID  uint    `gorm:"primaryKey"`
	P名称 string  `gorm:"column:name;uniqueIndex:udx_products_name"`
	P价格 float64 `gorm:"column:price;index:idx_products_price"`
}

func (*T产品) TableName() string { return "products" }

// T产品_漏写列名 的母语字段 S库存 没给 column，列名会落到非 ASCII。
type T产品_漏写列名 struct {
	ID  uint
	S库存 int // 无 column → 列名非 ASCII
}

func (*T产品_漏写列名) TableName() string { return "products" }

func TestNewT批量校验器(t *testing.T) {
	v批量校验器 := gormzhcn.NewT批量校验器([]any{&T产品{}}, gormzhcn.NewT配置项())
	require.NotNil(t, v批量校验器)
	require.NoError(t, v批量校验器.Validate())
}

func TestT批量校验器_Validate(t *testing.T) {
	err := gormzhcn.NewT批量校验器([]any{&T产品{}}, gormzhcn.NewT配置项()).Validate()
	require.NoError(t, err)
}

func TestT批量校验器_Validate_报错(t *testing.T) {
	err := gormzhcn.NewT批量校验器([]any{&T产品{}, &T产品_漏写列名{}}, gormzhcn.NewT配置项()).Validate()
	require.ErrorContains(t, err, "struct T产品_漏写列名 field S库存 column")
}
