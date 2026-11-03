package models

type T性别 string

const (
	M男性 T性别 = "男"
	F女性 T性别 = "女"
)

// T学生 mixes English fields and Chinese fields, and the cnm tags give each column a Chinese alias, so each generated column has a Chinese name: V班级编码, V名字, ...
//
// T学生 的字段英文、中文都有，cnm 标签给每一列起了中文别名，生成的列就都是中文：V班级编码、V名字……
type T学生 struct {
	ClassCode string `gorm:"column:class_code" cnm:"V班级编码"`
	V学号       int    `gorm:"column:student_num;primaryKey;autoIncrement:true" cnm:"V学号"`
	Name      string `gorm:"column:name" cnm:"V名字"`
	V性别       T性别    `gorm:"column:gender" cnm:"V性别"`
	V生日       string `gorm:"column:birthday" cnm:"V生日"`
}

func (*T学生) TableName() string {
	return "students"
}

// T班级 has Chinese fields alone, and gives each column an explicit database-safe ASCII name
//
// T班级 的字段全用中文，每一列都显式给了数据库安全的 ASCII 名字
type T班级 struct {
	V班级编码 string `gorm:"column:class_code;primaryKey"`
	V班级名称 string `gorm:"column:class_name;unique"`
	V班主任名 string `gorm:"column:head_teacher;unique"`
}

func (*T班级) TableName() string {
	return "classes"
}
