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
	"stvsljl.com/stvsl/influxdb"
)

const NodeTokenExpireDuration = time.Hour * 24 // 节点token过期时间

type NodeClaims struct {
	ID      string `json:"id"`
	Feature string `json:"feature"`
	jwt.StandardClaims
}

func nodeauthHandler(c *gin.Context) {
	// 绑定数据
	var claims NodeClaims
	err := c.ShouldBindJSON(&claims)
	if err != nil {
		CX001(c)
		return
	}
	// 校验数据
	n := Sql.NodeInformation{}
	str, err := n.GetFeatures(claims.ID)
	if err != nil {
		CX003(c)
		return
	}
	if str != claims.Feature {
		CX003(c)
		return
	}
	// 生成token
	token, _ := GenNodeToken(claims.ID, claims.Feature)
	c.JSON(http.StatusOK, gin.H{
		"code":    "CX200",
		"message": "success",
		"token":   token,
	})
}

func submitInfoHandler(c *gin.Context) {
	var submitInfo influxdb.SubmitInfo
	err := c.BindJSON(&submitInfo)
	submitInfo.NodeId = c.GetString("id")
	if err != nil {
		CX302(c)
		return
	}
	// 校验节点提交信息，若数据正常写入时序数据库
	if submitInfo.Check() {
		err := submitInfo.WriteToDB()
		if err != nil {
			CX411(c)
			return
		}
		// 向客户端返回校验信息
		c.JSON(http.StatusOK, gin.H{
			"code":    "CX200",
			"message": "success",
		})
		return
	} else {
		CX301(c)
	}
}

// 生成jwt
func GenNodeToken(id string, feature string) (string, error) {
	// 创建声明
	c := NodeClaims{
		id,
		feature,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(NodeTokenExpireDuration).Unix(),
			NotBefore: time.Now().Unix() - 10,
			Issuer:    "CSystemServer",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	return token.SignedString(RSA.RSA_PRIVATE_LOCAL)
}

// 解析jwt
func ParseNodeToken(tokenstr string) (*Claims, error) {
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

// 节点token校验
func NodeAuthMiddleware() func(c *gin.Context) {
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
