package models

import "time"

type CheckoutOrderPayment struct {
	ID            uint      `gorm:"primaryKey;autoIncrement;not null"`                     // 映射表唯一标识
	CheckoutID    uint      `gorm:"type:int unsigned;not null"`                            // 结算ID（关联 `checkout` 表）
	OrderID       uint      `gorm:"type:int unsigned;not null"`                            // 订单ID（关联 `orders` 表）
	PaymentID     uint      `gorm:"type:int unsigned;not null"`                            // 支付ID（关联 `payments` 表）
	Amount        float64   `gorm:"type:decimal(10,2);not null"`                           // 支付金额
	PaymentStatus string    `gorm:"type:enum('待支付', '已支付', '已取消');not null"`               // 支付状态
	CreatedAt     time.Time `gorm:"default:current_timestamp"`                             // 创建时间
	UpdatedAt     time.Time `gorm:"default:current_timestamp on update current_timestamp"` // 更新时间
}