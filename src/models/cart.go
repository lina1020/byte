package models

import "time"

type Cart struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null"`                     // 购物车唯一标识
	UserID    uint      `gorm:"type:int unsigned;not null"`                            // 【关联用户ID】
	CreatedAt time.Time `gorm:"default:current_timestamp"`                             // 购物车创建时间
	UpdatedAt time.Time `gorm:"default:current_timestamp on update current_timestamp"` // 购物车更新时间
}