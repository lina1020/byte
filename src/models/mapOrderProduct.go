package models

import "time"

type OrderProduct struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement;not null"`                     // 订单商品唯一标识
	OrderID            uint      `gorm:"type:int unsigned;not null"`                            // 订单ID
	ProductID          uint      `gorm:"type:int unsigned;not null"`                            // 商品ID
	Quantity           uint      `gorm:"type:int unsigned;not null"`                            // 商品数量
	PriceAtTimeOfOrder float64   `gorm:"type:decimal(10,2);not null"`                           // 下单时的商品价格
	CreatedAt          time.Time `gorm:"default:current_timestamp"`                             // 订单创建时间
	UpdatedAt          time.Time `gorm:"default:current_timestamp on update current_timestamp"` // 更新时间
	IsDelete           *int      `gorm:"type:TINYINT(2)"`                                       // 是否删除标志
}