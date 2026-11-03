package models

import "time"

// Example has Chinese fields without cnm tags, so the generated columns keep the field names: V名称, V性别, ...
//
// Example 的中文字段没写 cnm 标签，生成的列就沿用字段名：V名称、V性别……
type Example struct {
	ID        int32     `gorm:"column:id;primaryKey"`
	V名称       string    `gorm:"column:name;type:text;uniqueIndex:udx_examples_name"`
	V性别       string    `gorm:"column:gender"`
	V年龄       int       `gorm:"column:age"`
	V身高       int32     `gorm:"column:height"`
	V体重       int32     `gorm:"column:weight"`
	CreatedAt time.Time `gorm:"autoCreateTime;index"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
