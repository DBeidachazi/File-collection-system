// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWorklist = "worklist"

// Worklist mapped from table <worklist>
type Worklist struct {
	WorkID   int32  `gorm:"column:work_id;primaryKey;autoIncrement:true" json:"work_id"`
	CourseID int32  `gorm:"column:course_id;not null" json:"course_id"`
	Name     string `gorm:"column:name;not null" json:"name"`
	Path     string `gorm:"column:path;not null" json:"path"`
	Status   int32  `gorm:"column:status;not null" json:"status"` // 提交人数
}

// TableName Worklist's table name
func (*Worklist) TableName() string {
	return TableNameWorklist
}
