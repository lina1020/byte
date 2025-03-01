package models

import "time"

type CartProduct struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null"` // 购物车商品唯一标识
	CartID    uint      `gorm:"type:int unsigned;not null"`        // 关联的购物车ID
	ProductID uint      `gorm:"type:int unsigned;not null"`        // 关联的商品ID
	Quantity  uint      `gorm:"type:int unsigned;not null"`        // 商品数量
	CreatedAt time.Time `gorm:"default:current_timestamp"`         // 商品加入购物车时间
}