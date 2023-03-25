package utils

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

var jwtSceret = []byte("TodoList")

type Claims struct {
	Id  uint `json:"id"`
	jwt.StandardClaims
}

//签发用户token
func GenerateToken(id uint) (string, error) {
	nowTime := time.Now()
	expireTiime := nowTime.Add(24*time.Hour)
	claims := Claims{
		Id :   id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTiime.Unix(),
			Issuer: "todoList",
		},
	}
	tokenClaimd := jwt.NewWithClaims(jwt.SigningMethodES256,claims)
	token, err := tokenClaimd.SignedString(jwtSceret)
	return token, err
}

//验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{},func(token *jwt.Token) (i interface{}, e error){
		return jwtSceret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}