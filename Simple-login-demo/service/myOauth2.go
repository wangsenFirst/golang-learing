package service

import (
	"Simple-login-demo/dao"
	"Simple-login-demo/models"
	wechat "Simple-login-demo/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 账号密码授权
func ImplicitGrant(login *models.Login) map[string]interface{} {
	// 根据账号查询数据库
	// 查询时会自动选择 `id`, `name` 字段
	queryUser, err := dao.QueryUser(login)
	if err != nil {
		println("ImplicitGrant查询：", err)
		return nil
	}
	fmt.Println("查询到的用户：", queryUser)
	// SELECT `id`, `name` FROM `users` LIMIT 10

	// 判断用户状态

	// JWT生成token
	token, err := GenerateToken(&queryUser)
	if err != nil {
		fmt.Println("生成JWT失败ERRO：", err)
		return nil
	}
	// 保存到redis生成白名单

	// 构造返回数据
	data := make(map[string]interface{}, 4)
	data["access_token"] = token
	data["expires_in"] = queryUser.ExpiresAt
	data["user_name"] = queryUser.Name
	if err != nil {
		fmt.Println("解析JWT失败ERRO：", err)
	}
	return data
}

// code 登入
func WeChatCodeLogin(code string) map[string]interface{} {
	response, err := wechat.GetUserIdByCode(code)
	if err != nil {
		data := make(map[string]interface{}, 4)
		data["erro"] = response
		return data
	}
	//使用UserId登入
	queryUser, err := dao.QueryUserByUserId(response.UserId)
	if err != nil {
		fmt.Println("未查询到用户", err)
		return nil
	}
	// JWT生成token
	token, err := GenerateToken(&queryUser)
	if err != nil {
		fmt.Println("生成JWT失败ERRO：", err)
		return nil
	}
	// 保存到redis生成白名单

	// 构造返回数据
	data := make(map[string]interface{}, 4)
	data["user"] = token
	data["expires_in"] = queryUser.ExpiresAt
	data["user_name"] = queryUser.Name
	if err != nil {
		fmt.Println("解析JWT失败ERRO：", err)
	}
	return data

}

// 客户端授权

// code授权（先要登入获取code）

// 校验密码
func CheckPassword(password string, user models.User) (bool, error) {
	// TODO 密码解密

	// 密码校验

	return true, nil
}

// 生成token
func GenerateToken(user *models.User) (string, error) {
	// 过期时间
	nowTime := time.Now()
	expireTime := nowTime.Add(300 * time.Second)
	issuer := "frank"

	user.ExpiresAt = expireTime.Unix()
	user.Issuer = issuer

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, user).SignedString([]byte("golang"))
	return token, err
}

// 解析token
func ParseToken(token string) (*models.User, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &models.User{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*models.User); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
