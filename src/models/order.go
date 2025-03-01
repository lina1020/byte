package models

import "time"

type Order struct {
	ID            uint      `gorm:"primaryKey;autoIncrement;not null"`                     // 订单唯一标识
	TotalPrice    float64   `gorm:"type:decimal(10,2);not null"`                           // 订单总价
	Status        string    `gorm:"type:enum('待支付', '已支付', '已取消');not null"`               // 订单状态
	Currency      string    `gorm:"type:varchar(10);not null"`                             // 存储货币类型
	StreetAddress string    `gorm:"type:varchar(30);not null"`                             // 街道地址
	City          string    `gorm:"type:varchar(30);not null"`                             // 城市
	State         string    `gorm:"type:varchar(30)"`                                      // 州/省
	Country       string    `gorm:"type:varchar(30);not null"`                             // 国家
	ZipCode       string    `gorm:"type:varchar(30);not null"`                             // 邮政编码
	CreatedAt     time.Time `gorm:"default:current_timestamp"`                             // 订单创建时间
	UpdatedAt     time.Time `gorm:"default:current_timestamp on update current_timestamp"` // 订单更新时间
	Operator      string    `gorm:"type:varchar(10);not null"`                             // 操作员
	Creator       string    `gorm:"type:varchar(10);not null"`                             // 创建者
	IsDelete      *int      `gorm:"type:TINYINT(2)"`                                       // 是否删除标志
}