package models

type T性别 string

const (
	M男性 T性别 = "男"
	F女性 T性别 = "女"
)

/*
假设开发者英语水平不高，只会些塑料英语，甚至有字段名对齐的强迫症，则允许开发者使用别名
*/

type T学生 struct {
	ClassCode string `gorm:"column:class_code" cnm:"V班级编码"`
	V学号       int    `gorm:"column:V_665B_F753;primaryKey;autoIncrement:true" cnm:"V学号" mom:"mcp:S63;"`
	Name      string `gorm:"column:name" cnm:"V名字"`
	V性别       T性别    `gorm:"column:V_2760_2B52;" cnm:"V性别" mom:"mcp:S63;"`
	V生日       string `gorm:"column:V_1F75_E565;" cnm:"V生日" mom:"mcp:S63;"`
}

func (*T学生) TableName() string {
	return "students"
}

/*
彻底放弃治疗直接使用汉语编程，毕竟汉语天然的表意比较全面，而且非常精悍，而且字段名长度相同有利于代码对齐和整洁（这点对强迫症非常友善）
*/

type T班级 struct {
	V班级编码 string `gorm:"column:class_code;primaryKey;"`
	V班级名称 string `gorm:"column:class_name;unique"`
	V班主任名 string `gorm:"column:V_ED73_3B4E_FB4E_0D54;unique;" mom:"mcp:S63;"`
}

func (*T班级) TableName() string {
	return "classes"
}
