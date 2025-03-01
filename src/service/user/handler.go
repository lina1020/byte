package main

import (
	"byteSmart/src/models"
	"byteSmart/src/rpc/user"
	"byteSmart/src/storage/database"
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// UserServiceImpl 用户服务实现
type UserServiceImpl struct {
	user.UserServiceServer
}

// GetUserExistInformation 获取用户是否存在
func (s *UserServiceImpl) GetUserExistInformation(ctx context.Context, request *user.UserExistReq) (resp *user.UserExistResp, err error) {
	// 从数据库中查询用户是否存在
	err = database.Client.Where("email = ?", request.Email).First(&models.User{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp = &user.UserExistResp{
				StatusCode: 0,
				StatusMsg:  "用户不存在",
				Existed:    false,
			}
			return
		} else {
			resp = &user.UserExistResp{
				StatusCode: 1,
				StatusMsg:  "查询失败",
				Existed:    false,
			}
			err = status.Errorf(http.StatusInternalServerError, "查询失败")
			return
		}
	}
	resp = &user.UserExistResp{
		StatusCode: 0,
		StatusMsg:  "用户存在",
		Existed:    true,
	}
	return
}

// Register 注册
func (s *UserServiceImpl) Register(ctx context.Context, request *user.RegisterReq) (resp *user.RegisterResp, err error) {
	exist, err := s.GetUserExistInformation(ctx, &user.UserExistReq{Email: request.Email})
	if  err != nil {
		err = status.Errorf(http.StatusInternalServerError, "查询失败")
	} else if exist.Existed {
		resp = &user.RegisterResp{
			StatusCode: 1,
			StatusMsg:  "用户已存在",
		}
		return
	} else if request.Password != request.ConfirmPassword {
		resp = &user.RegisterResp{
			StatusCode: 2,
			StatusMsg:  "两次密码不一致",
		}
		return
	} else {
		hashedPassword, e1 := HashPassword(ctx, request.Password)
		err = e1
		if err != nil {
			resp = &user.RegisterResp{
				StatusCode: 3,
				StatusMsg:  "密码哈希失败",
			}
			return
		}
		newUser := models.User{
			Email:    request.Email,
			Password: hashedPassword,
		}
		if err = database.Client.AutoMigrate(&newUser); err != nil {
			resp = &user.RegisterResp{
				StatusCode: 3,
				StatusMsg:  "数据库迁移失败",
			}
			return
		}
		token, e2 := GenerateJWT(newUser.ID)
		err = e2
		if err != nil {
			resp = &user.RegisterResp{
				StatusCode: 4,
				StatusMsg:  "生成JWT失败",
			}
			return
		}
		resp = &user.RegisterResp{
			StatusCode: 0,
			StatusMsg:  "注册成功",
			UserId:     newUser.ID,
			Token:      token,
		}
		return
	}
	return nil, nil
}

// Login 登录
func (s *UserServiceImpl) Login(ctx context.Context, request *user.LoginReq) (resp *user.LoginResp, err error) {
	var currentUser models.User
	result := database.Client.Where("email = ?", request.Email).First(&currentUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			resp = &user.LoginResp{
				StatusCode: 1,
				StatusMsg:  "用户不存在",
			}
		} else {
			resp = &user.LoginResp{
				StatusCode: 2,
				StatusMsg:  "查询失败",
			}
		}
		return
	}
	if !CheckPassword(request.Password, currentUser.Password) {
		resp = &user.LoginResp{
			StatusCode: 3,
			StatusMsg:  "密码错误",
		}
		return
	}
	token, err := GenerateJWT(currentUser.ID)
	if err != nil {
		resp = &user.LoginResp{
			StatusCode: 4,
			StatusMsg:  "生成JWT失败",
		}
		return
	}
	resp = &user.LoginResp{
		StatusCode: 0,
		StatusMsg:  "登录成功",
		Token:      token,
	}
	return
}

// HashPassword 获取密码哈希
func HashPassword(ctx context.Context, password string) (string, error) {
	// cost 表示计算复杂度，越大越安全，但是也越耗时
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// CheckPassword 检查密码是否正确
func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// GenerateJWT 生成JWT
func GenerateJWT(userID uint32) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,                                // 用户ID
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // 过期时间
	})
	signedToken, err := token.SignedString([]byte("secret")) // 签名
	if err != nil {
		return "", err
	}
	return "Bearer " + signedToken, nil
}

// ParseJWT 验证token
func ParseJWT(tokenString string) (uint32, error) {
	// 去除Bearer 前缀
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	// 解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("签名方法错误")
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return 0, err
	}
	// 验证token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["userID"].(float64)
		if !ok {
			return 0, errors.New("用户ID无效")
		}
		return uint32(userID), nil
	} else {
		return 0, errors.New("token无效")
	}
}
