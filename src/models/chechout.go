package models

import "time"

type Checkout struct {
	ID             uint      `gorm:"primaryKey;autoIncrement;not null"`                     // 结算唯一标识
	UserID         uint      `gorm:"type:int unsigned;not null"`                            // 用户ID
	FirstName      string    `gorm:"type:varchar(20);not null"`                             // 用户名（结算时的姓名）
	LastName       string    `gorm:"type:varchar(20);not null"`                             // 用户姓氏
	Email          string    `gorm:"type:varchar(30);not null"`                             // 用户邮箱
	CreatedAt      time.Time `gorm:"default:current_timestamp"`                             // 创建时间
	UpdatedAt      time.Time `gorm:"default:current_timestamp on update current_timestamp"` // 更新时间
	CheckoutStatus string    `gorm:"type:enum('待结算', '已结算', '已取消');default:'待结算'"`          // 结算状态
}