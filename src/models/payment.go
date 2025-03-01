package models

import "time"

type Payment struct {
	ID                        uint      `gorm:"primaryKey;autoIncrement;not null"`                     // 支付唯一标识
	Amount                    float64   `gorm:"type:decimal(10,2);not null"`                           // 支付金额
	Status                    string    `gorm:"type:enum('待支付', '已支付', '已取消');not null"`               // 支付状态
	TransactionID             string    `gorm:"type:varchar(100);not null"`                            // 存储交易ID
	CreditCardNumber          string    `gorm:"type:varchar(20)"`                                      // 信用卡号
	CreditCardCVV             uint      `gorm:"type:int(3) unsigned"`                                  // 信用卡 CVV
	CreditCardExpirationYear  uint      `gorm:"type:int(4) unsigned"`                                  // 信用卡过期年份
	CreditCardExpirationMonth uint      `gorm:"type:int(2) unsigned"`                                  // 信用卡过期月份
	CreatedAt                 time.Time `gorm:"default:current_timestamp"`                             // 支付创建时间
	UpdatedAt                 time.Time `gorm:"default:current_timestamp on update current_timestamp"` // 支付更新时间
}