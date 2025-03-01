package models

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;not null"`                     // 商品唯一标识
	Name        string    `gorm:"type:varchar(20);not null"`                             // 商品名称
	Description string    `gorm:"type:json"`                                             // 商品描述
	Price       string    `gorm:"type:varchar(10);not null"`                             // 商品价格
	Stock       uint      `gorm:"type:int unsigned;not null"`                            // 商品库存
	IsDelete    *int      `gorm:"type:TINYINT(2)"`                                       // 是否删除标志
	Operator    string    `gorm:"type:varchar(10);not null"`                             // 操作员
	Creator     string    `gorm:"type:varchar(10);not null"`                             // 创建者
	CreatedAt   time.Time `gorm:"default:current_timestamp"`                             // 商品创建时间
	UpdatedAt   time.Time `gorm:"default:current_timestamp on update current_timestamp"` // 商品信息更新时间
}