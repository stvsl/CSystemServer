package Service

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"stvsljl.com/stvsl/AES"
	"stvsljl.com/stvsl/RSA"
	"stvsljl.com/stvsl/Sql"
)

const TokenExpireDuration = time.Hour * 2 // token过期时间

type Claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

// 生成jwt
func GenToken(id string) (string, error) {
	// 创建声明
	c := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			NotBefore: time.Now().Unix() - 10,
			Issuer:    "CSystemServer",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	priv, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(RSA.RSA_PRIVATE_LOCAL))
	return token.SignedString(priv)
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
	var claims struct {
		ID       string `json:"id"`
		Feature  string `json:"feature"`
		Fragment string `json:"passwd"`
	}
	err := c.ShouldBindJSON(&claims)
	if err != nil {
		CX001(c)
		return
	}
	// base64解码
	deid, _ := base64.StdEncoding.DecodeString(claims.ID)
	defragment, _ := base64.StdEncoding.DecodeString(claims.Fragment)
	defeature, _ := base64.StdEncoding.DecodeString(claims.Feature)
	// RSA解密
	id, err := RSA.Decrypt(deid, []byte(RSA.RSA_PRIVATE_LOCAL))
	if err != nil {
		CX304(c)
		return
	}
	feature, err := RSA.Decrypt(defeature, []byte(RSA.RSA_PRIVATE_LOCAL))
	if err != nil {
		CX304(c)
		return
	}
	fragment, err := RSA.Decrypt(defragment, []byte(RSA.RSA_PRIVATE_LOCAL))
	if err != nil {
		CX304(c)
		return
	}
	idstr := string(id)
	featurestr := string(feature)
	fragmentstr := string(fragment)
	// 校验节点信息
	// 安全性校验-帐号存在性校验
	n := Sql.AccountInformations{}
	b, err := n.Exist(idstr)
	if err != nil {
		CX401(c)
		return
	}
	if !b {
		CX003(c)
		return
	}
	// 安全性校验-密码特征校验
	s, err := n.GetPasswdFragment(idstr)
	if err != nil {
		CX401(c)
		return
	}
	if s != fragmentstr {
		CX003(c)
		return
	}
	// 签发token
	tokens, _ := GenToken(idstr)
	// 从loginings中删除暂存记录
	l := Sql.Logining{}
	aes, err := l.Read(featurestr) // 导出暂存记录
	if err != nil {
		CX401(c)
		return
	}
	// token使用AES加密
	token, _ := AES.AesEncrypt([]byte(tokens), []byte(aes))
	// base64编码
	tokenstr := base64.StdEncoding.EncodeToString(token)
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX200",
		"message": "success",
		"token":   tokenstr,
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
		c.Next()
	}
}
