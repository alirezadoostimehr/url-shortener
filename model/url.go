package model

import "time"

type Url struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Srt       string    `gorm:"column:srt;not null;size:255" json:"short"`
	Lng       string    `gorm:"column:lng;not null;size:255" json:"long"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
