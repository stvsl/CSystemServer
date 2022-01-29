package Service

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"stvsljl.com/stvsl/RSA"
	"stvsljl.com/stvsl/Sql"
)

const TokenExpireDuration = time.Hour * 2 // token过期时间

type Claims struct {
	ID       string `json:"id"`
	Fragment string `json:"fragment"`
	jwt.StandardClaims
}

// 生成jwt
func GenToken(id string, fragment string) (string, error) {
	// 创建声明
	c := Claims{
		id,
		fragment,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			NotBefore: time.Now().Unix() - 10,
			Issuer:    "CSystemServer",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	return token.SignedString(RSA.RSA_PRIVATE_LOCAL)
}

// 解析jwt
func ParseToken(tokenstr string) (*Claims, error) {
	//  解析token
	token, err := jwt.ParseWithClaims(tokenstr, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return RSA.RSA_PUBLIC_LOCAL, nil
	})
	if err != nil {
		return nil, err
	}
	// 校验token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token异常")
}

// 签发token密钥
func authHandler(c *gin.Context) {
	var claims Claims
	err := c.ShouldBindJSON(&claims)
	if err != nil {
		CX001(c)
		return
	}
	// 校验节点信息
	// 安全性校验-帐号存在性校验
	n := Sql.AccountInformations{}
	b, err := n.Exist(claims.ID)
	if err != nil {
		CX401(c)
		return
	}
	if !b {
		CX003(c)
		return
	}
	// 安全性校验-密码残片校验
	s, err := n.GetFragment(claims.ID)
	if err != nil {
		CX401(c)
		return
	}
	if s != claims.Fragment {
		CX003(c)
		return
	}
	// 签发token
	tokens, _ := GenToken(claims.ID, claims.Fragment)
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX200",
		"message": "success",
		"token":   tokens,
	})

}

// token校验
func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			CX002(c)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "sintok") {
			CX006(c)
			c.Abort()
			return
		}
		// 解析token部分
		message, err := ParseToken(parts[1])
		if err != nil {
			CX007(c)
			c.Abort()
			return
		}
		// 将当前请求的信息保存到请求的上下文c上
		c.Set("id", message.ID)
		c.Set("fragment", message.Fragment)
		c.Next()
	}
}
