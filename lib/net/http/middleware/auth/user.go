package auth

import (
	"encoding/json"
	"errors"
	"fmt"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const (
	USER_TYPE_ADMIN = iota
	USER_TYPE_MEMBER
)

type User struct {
	ID        uint   `json:"id"`
	TokenSalt string `json:"tokenSalt"`
}

func GetCurrentUser(c *gin.Context) (*User, error) {
	jwtClaims := jwt.ExtractClaims(c)
	fmt.Println(jwtClaims)
	if jwtClaims == nil {
		return nil, errors.New("无法解析用户信息")
	}

	js, err := json.Marshal(jwtClaims)
	if err != nil {
		return nil, err
	}

	user := &User{}
	err = json.Unmarshal(js, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
