package cart

import (
	"byteSmart/src/constant/config"
	"byteSmart/src/constant/strings"
	"byteSmart/src/extra/tracing"
	"byteSmart/src/rpc/cart"
	grpc2 "byteSmart/src/utils/grpc"
	"byteSmart/src/utils/logging"
	"byteSmart/src/web/models"
	"byteSmart/src/web/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Client cart.CartServiceClient

// init 是 Go 语言中的特殊函数，会在包被初始化时自动执行。
// 此函数用于初始化 CartService 的 gRPC 客户端。
func init() {
	// 调用 grpc2.Connect 函数，传入 Cart RPC 服务器的名称，以建立与 Cart RPC 服务器的连接。
	// config.CartRpcServerName 是从配置文件中获取的 Cart RPC 服务器的名称。
	conn := grpc2.Connect(config.CartRpcServerName)
	// 使用建立好的连接 conn 创建一个新的 CartService 客户端实例。
	// cart.NewCartServiceClient 是由 gRPC 生成的函数，用于创建 CartService 的客户端。
	Client = cart.NewCartServiceClient(conn)
}

// rpc CreateCart(CreateCartReq) returns (CreateCartResp) {}
// rpc AddItem(AddItemReq) returns (AddItemResp) {}
// rpc GetCart(GetCartReq) returns (GetCartResp) {}
// rpc EmptyCart(EmptyCartReq) returns (EmptyCartResp) {}
// AddItemHandler 处理向购物车添加商品项的 HTTP 请求。
// 该函数接收一个 *gin.Context 类型的参数，用于处理 HTTP 请求和响应。
func AddItemHandler(c *gin.Context) {
	// 声明一个 models.AddItemReq 类型的变量 req，用于存储从请求中解析的参数。
	var req models.AddItemReq
	// 使用 tracing.Tracer 启动一个新的追踪 span，用于记录处理该请求的性能信息。
	// 该 span 的名称为 "AddItemHandler"，并在函数结束时自动结束该 span。
	_, span := tracing.Tracer.Start(c.Request.Context(), "AddItemHandler")
	defer span.End()
	// 设置追踪 span 的主机名信息，方便后续的性能分析和故障排查。
	logging.SetSpanWithHostname(span)
	// 创建一个日志记录器，用于记录处理该请求过程中的日志信息。
	// 日志记录器的名称为 "GateWay.AddItem"，并关联当前请求的上下文。
	logger := logging.LogService("GateWay.AddItem").WithContext(c.Request.Context())
	// 尝试从请求的查询参数中解析出 req 结构体所需的字段。
	// 如果解析失败，说明请求参数格式不正确，返回一个包含错误信息的 JSON 响应。
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, models.AddItemRes{
			StatusCode: strings.GateWayParamsErrorCode,
			StatusMsg:  strings.GateWayParamsError,
		})
		return
	}
	// 声明一个 *cart.AddItemResp 类型的变量 res，用于存储 gRPC 服务的响应结果。
	// 声明一个 error 类型的变量 err，用于存储可能出现的错误信息。
	var res *cart.AddItemResp
	var err error
	// 调用 CartService 的 AddItem 方法，向购物车中添加商品项。
	// 传入当前请求的上下文和一个 *cart.AddItemReq 类型的请求对象。
	// 请求对象中包含用户 ID 和要添加的商品项信息。
	res, err = Client.AddItem(c.Request.Context(), &cart.AddItemReq{
		UserId: uint32(req.UserId),
		Item: &cart.CartItem{
			ProductId: req.Item.ProductId,
			Quantity:  req.Item.Quantity,
		},
	})
	// 如果调用 gRPC 服务时出现错误，记录警告日志并返回包含错误信息的响应。
	if err != nil {
		logger.WithFields(logrus.Fields{
			"UserId":       req.UserId,
			"ItemID":       req.Item.ProductId,
			"ItemQuantity": req.Item.Quantity,
		}).Warnf("Error when trying to connect with AddItemHandler")
		c.Render(http.StatusOK, utils.CustomJSON{Data: res, Context: c})
		return
	}
	// 如果调用 gRPC 服务成功，记录成功日志并返回包含响应结果的响应。
	logger.WithFields(logrus.Fields{
		"UserId":       req.UserId,
		"ItemID":       req.Item.ProductId,
		"ItemQuantity": req.Item.Quantity,
	}).Infof("Add Items success")
	c.Render(http.StatusOK, utils.CustomJSON{Data: res, Context: c})
}
func CreateCartHandler(c *gin.Context) {

	var req models.CreateCartReq

	_, span := tracing.Tracer.Start(c.Request.Context(), "CreateCartHandler")
	defer span.End()
	logging.SetSpanWithHostname(span)

	logger := logging.LogService("GateWay.CreateCart").WithContext(c.Request.Context())
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, models.CreateCartRes{
			StatusCode: strings.GateWayParamsErrorCode,
			StatusMsg:  strings.GateWayParamsError,
		})
		return
	}

	var res *cart.CreateCartResp
	var err error
	res, err = Client.CreateCart(c.Request.Context(), &cart.CreateCartReq{
		UserId: uint32(req.UserId),
	})

	if err != nil {
		logger.WithFields(logrus.Fields{
			"UserId": req.UserId,
		}).Warnf("Error when trying to connect with CreateCartHandler")
		c.Render(http.StatusOK, utils.CustomJSON{Data: res, Context: c})
		return
	}

	logger.WithFields(logrus.Fields{
		"UserId": req.UserId,
	}).Infof("Create cart success")

	c.Render(http.StatusOK, utils.CustomJSON{Data: res, Context: c})
}

func GetCartHandler(c *gin.Context) {

	var req models.GetCartReq

	_, span := tracing.Tracer.Start(c.Request.Context(), "GetCartHandler")
	defer span.End()
	logging.SetSpanWithHostname(span)

	logger := logging.LogService("GateWay.GetCart").WithContext(c.Request.Context())
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, models.GetCartRes{
			StatusCode: strings.GateWayParamsErrorCode,
			StatusMsg:  strings.GateWayParamsError,
		})
		return
	}

	var res *cart.GetCartResp
	var err error
	res, err = Client.GetCart(c.Request.Context(), &cart.GetCartReq{
		UserId: uint32(req.UserId),
	})

	if err != nil {
		logger.WithFields(logrus.Fields{
			"UserId": req.UserId,
		}).Warnf("Error when trying to connect with GetCartHandler")
		c.Render(http.StatusOK, utils.CustomJSON{Data: res, Context: c})
		return
	}

	logger.WithFields(logrus.Fields{
		"UserId": req.UserId,
	}).Infof("Get cart success")

	c.Render(http.StatusOK, utils.CustomJSON{Data: res, Context: c})
}

func EmptyCartHandler(c *gin.Context) {
	var req models.EmptyCartReq

	_, span := tracing.Tracer.Start(c.Request.Context(), "EmptyCartHandler")
	defer span.End()
	logging.SetSpanWithHostname(span)

	logger := logging.LogService("GateWay.EmptyCart").WithContext(c.Request.Context())
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, models.EmptyCartRes{
			StatusCode: strings.GateWayParamsErrorCode,
			StatusMsg:  strings.GateWayParamsError,
		})
		return
	}

	var res *cart.EmptyCartResp
	var err error
	res, err = Client.EmptyCart(c.Request.Context(), &cart.EmptyCartReq{
		UserId: uint32(req.UserId),
	})

	if err != nil {
		logger.WithFields(logrus.Fields{
			"UserId": req.UserId,
		}).Warnf("Error when trying to connect with EmptyCartHandler")
		c.Render(http.StatusOK, utils.CustomJSON{Data: res, Context: c})
		return
	}

	logger.WithFields(logrus.Fields{
		"UserId": req.UserId,
	}).Infof("Empty cart success")

	c.Render(http.StatusOK, utils.CustomJSON{Data: res, Context: c})
}
