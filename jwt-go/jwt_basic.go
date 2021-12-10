package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func main() {
	mySigningKey := []byte("woshiaaronmegsfangxing")
	myClaims := MyClaims{
		Username: "aaaronmegs",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,          //60s前开始生效
			ExpiresAt: time.Now().Unix() * 60 * 60 * 2, // 过期时间2小时
			Issuer:    "aaronmegs",                     // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	// 加密
	s, e := token.SignedString(mySigningKey)
	if e != nil {
		fmt.Printf("%s", e)
	}
	fmt.Println(s)

	// 校验 - 解密
	token, err := jwt.ParseWithClaims(s, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println(token.Claims.(*MyClaims).Username)

}
