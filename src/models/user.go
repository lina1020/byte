package models

import "time"

type User struct {
	ID       uint32 `gorm:"primaryKey;autoIncrement;not null"` // 用户唯一标识
	Name     string `gorm:"type:varchar(20);not null"`         // 用户名
	Slug     string `gorm:"type:varchar(20);not null;unique"`  // 唯一标识
	Password string `gorm:"type:varchar(100);not null;unique"` // 密码
	Email    string `gorm:"type:varchar(30);not null"`         // 用户邮箱
	// Avatar    string    `redis:"Avatar"`                                               // 用户头像
	CreatedAt time.Time `gorm:"default:current_timestamp"`                             // 用户创建时间
	UpdatedAt time.Time `gorm:"default:current_timestamp on update current_timestamp"` // 用户信息更新时间
	IsDeleted *int      `gorm:"type:TINYINT(2)"`                                       // 是否删除标志
}
