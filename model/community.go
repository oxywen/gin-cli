package model

import "time"

type Community struct {
	ID            int64
	CommunityId   int64     `gorm:"column:community_id"`
	CommunityName string    `gorm:"column:community_name"`
	Introduction  string    `gorm:"column:introduction"`
	CreateTime    time.Time `gorm:"column:create_time"`
	UpdateTime    time.Time `gorm:"column:update_time"`
}
