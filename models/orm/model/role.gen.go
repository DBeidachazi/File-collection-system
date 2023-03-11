// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameRole = "role"

// Role mapped from table <role>
type Role struct {
	StuID    int32  `gorm:"column:stu_id;not null" json:"stu_id"`
	RoleID   int32  `gorm:"column:role_id;primaryKey;autoIncrement:true" json:"role_id"`
	RoleName string `gorm:"column:role_name;not null" json:"role_name"`
	ClassID  int32  `gorm:"column:class_id;not null" json:"class_id"`
	Status   int32  `gorm:"column:status;not null" json:"status"` // 1:已提交 2:未提交 3:发布者(不显示)
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}
