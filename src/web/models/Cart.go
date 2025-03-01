package models

import "byteSmart/src/rpc/cart"

// CreateCartReq 创建购物车请求结构体
type CreateCartReq struct {
	Token   string `form:"token" binding:"required"` // 用户的身份令牌
	UserId  uint32 `form:"user_id" binding:"required"` // 用户ID
}

// CreateCartRes 创建购物车响应结构体
type CreateCartRes struct {
	StatusCode int    `json:"status_code"` // 状态码
	StatusMsg  string `json:"status_msg"`  // 状态消息
	Cart       cart.Cart   `json:"cart"`        // 返回的购物车对象
}

// GetCartReq 获取购物车请求结构体
type GetCartReq struct {
	Token   string `form:"token" binding:"required"` // 用户的身份令牌
	UserId  uint32 `form:"user_id" binding:"required"` // 用户ID
}

// GetCartRes 获取购物车响应结构体
type GetCartRes struct {
	StatusCode int    `json:"status_code"` // 状态码
	StatusMsg  string `json:"status_msg"`  // 状态消息
	Cart       cart.Cart   `json:"cart"`        // 返回的购物车对象
}

// AddItemReq 添加商品到购物车请求结构体
type AddItemReq struct {
	Token    string   `form:"token" binding:"required"`   // 用户的身份令牌
	UserId   uint32   `form:"user_id" binding:"required"` // 用户ID
	Item     cart.CartItem `form:"item" binding:"required"`    // 要添加的商品信息

}

// AddItemRes 添加商品到购物车响应结构体
type AddItemRes struct {
	StatusCode int    `json:"status_code"` // 状态码
	StatusMsg  string `json:"status_msg"`  // 状态消息
}

// EmptyCartReq 清空购物车请求结构体
type EmptyCartReq struct {
	Token   string `form:"token" binding:"required"` // 用户的身份令牌
	UserId  uint32 `form:"user_id" binding:"required"` // 用户ID
}

// EmptyCartRes 清空购物车响应结构体
type EmptyCartRes struct {
	StatusCode int    `json:"status_code"` // 状态码
	StatusMsg  string `json:"status_msg"`  // 状态消息
}
