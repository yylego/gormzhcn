// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yylego/gormcngen

//go:build !gormcngen_generate

// Generated from: gormcnm_test.go:50 -> models.TestGenGormMomAndCnm.func2
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import "github.com/yylego/gormcnm"

func (T *T学生) Columns() *T学生Columns {
	return &T学生Columns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		V班级编码: gormcnm.Cnm(T.ClassCode, "class_code"),
		V学号:   gormcnm.Cnm(T.V学号, "V_665B_F753"),
		V名字:   gormcnm.Cnm(T.Name, "name"),
		V性别:   gormcnm.Cnm(T.V性别, "V_2760_2B52"),
		V生日:   gormcnm.Cnm(T.V生日, "V_1F75_E565"),
	}
}

type T学生Columns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	V班级编码 gormcnm.ColumnName[string]
	V学号   gormcnm.ColumnName[int]
	V名字   gormcnm.ColumnName[string]
	V性别   gormcnm.ColumnName[T性别]
	V生日   gormcnm.ColumnName[string]
}

func (T *T班级) Columns() *T班级Columns {
	return &T班级Columns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		V班级编码: gormcnm.Cnm(T.V班级编码, "class_code"),
		V班级名称: gormcnm.Cnm(T.V班级名称, "class_name"),
		V班主任名: gormcnm.Cnm(T.V班主任名, "V_ED73_3B4E_FB4E_0D54"),
	}
}

type T班级Columns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	V班级编码 gormcnm.ColumnName[string]
	V班级名称 gormcnm.ColumnName[string]
	V班主任名 gormcnm.ColumnName[string]
}
