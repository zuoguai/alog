package utils

import (
	"alog/internal/pkg"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

func NewUuid() string {
	uuid_ := uuid.New()
	return uuid_.String()
}

type JWTClaims struct {
	UserID         uint   `json:"user_id"`
	Username       string `json:"username"`
	StandardClaims jwt.StandardClaims
}

func (j JWTClaims) Valid() error {
	return nil
}

// GenerateJwtToken 生成token
func GenerateJwtToken(secret string, issuer string, expire int64, userID uint, username string) (string, error) {
	if secret == "" {
		secret = pkg.JWTDefaultSecret
	}
	if issuer == "" {
		issuer = pkg.JWTDefaultIssuer
	}
	if expire == 0 {
		expire = pkg.TokenExpireDuration
	}
	hmacSampleSecret := []byte(secret) //密钥，不能泄露
	token := jwt.New(jwt.SigningMethodHS256)
	nowTime := time.Now().Unix()
	token.Claims = JWTClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: nowTime,                                                    // 签名生效时间
			ExpiresAt: time.Now().Add(time.Duration(expire) * time.Second).Unix(), // 签名过期时间
			Issuer:    issuer,                                                     // 签名颁发者
		},
	}
	tokenString, err := token.SignedString(hmacSampleSecret)
	return tokenString, err
}

// ParseJwtToken 解析token
func ParseJwtToken(tokenString string, secret string) (*JWTClaims, error) {
	if secret == "" {
		secret = pkg.JWTDefaultSecret
	}
	var hmacSampleSecret = []byte(secret)
	//前面例子生成的token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	claims := token.Claims.(*JWTClaims)
	return claims, nil
}
