// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameClass = "class"

// Class mapped from table <class>
type Class struct {
	ClassID   int32  `gorm:"column:class_id;primaryKey;autoIncrement:true" json:"class_id"`
	Name      string `gorm:"column:name;not null" json:"name"`
	ClassName string `gorm:"column:class_name;not null" json:"class_name"`
}

// TableName Class's table name
func (*Class) TableName() string {
	return TableNameClass
}
