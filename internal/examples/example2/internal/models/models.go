package models

import "time"

type Example struct {
	ID        int32     `gorm:"column:id;primaryKey;" json:"id"`
	V名称       string    `gorm:"column:name;type:text;uniqueIndex:udx_examples_name"`
	V性别       string    `gorm:"column:gender;"`
	V年龄       int       `gorm:"column:age;" json:"age"` //理论上不要直接给model添加json标签，因为那是view层的逻辑，但实际上假如非这样做也能处理
	V身高       int32     `gorm:"column:height;"`
	V体重       int32     `gorm:"column:weight;"`
	CreatedAt time.Time `gorm:"autoCreateTime;index;"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;uniqueIndex;"`
}
