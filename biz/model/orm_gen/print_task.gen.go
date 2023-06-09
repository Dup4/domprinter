// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm_gen

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePrintTask = "print_task"

// PrintTask mapped from table <print_task>
type PrintTask struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`  // Print Task ID
	State      string         `gorm:"column:state;not null;default:Unknown" json:"state"` // Print Task State
	UserName   string         `gorm:"column:user_name;not null" json:"user_name"`         // User Name
	TeamName   string         `gorm:"column:team_name;not null" json:"team_name"`         // Team Name
	TeamID     string         `gorm:"column:team_id;not null" json:"team_id"`             // Team ID
	Location   string         `gorm:"column:location;not null" json:"location"`           // Location
	Language   string         `gorm:"column:language;not null" json:"language"`           // Language
	FileName   string         `gorm:"column:file_name;not null" json:"file_name"`         // File Name
	SourceCode []byte         `gorm:"column:source_code;not null" json:"source_code"`     // Source Code
	SubmitTime time.Time      `gorm:"column:submit_time;not null;default:current_timestamp(3)" json:"submit_time"`
	CreatedAt  time.Time      `gorm:"column:created_at;not null;default:current_timestamp()" json:"created_at"` // Print task create time
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null;default:current_timestamp()" json:"updated_at"` // Print task update time
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                                      // Print task delete time
}

// TableName PrintTask's table name
func (*PrintTask) TableName() string {
	return TableNamePrintTask
}
