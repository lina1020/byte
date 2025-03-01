package models

import "time"

type UserOrder struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null"` // 映射表唯一标识
	UserID    uint      `gorm:"type:int unsigned;not null"`        // 用户ID
	OrderID   uint      `gorm:"type:int unsigned;not null"`        // 订单ID
	CreatedAt time.Time `gorm:"default:current_timestamp"`         // 创建时间
}