package util

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetTokenSession(c *gin.Context, token string) error {
	session := sessions.Default(c)
	session.Set("token", token)
	return session.Save()
}

func GetTokenSession(c *gin.Context) (string, error) {
	session := sessions.Default(c)
	token := session.Get("token")
	if token == nil {
		return "", errors.New("cannot get token from session")
	}

	return token.(string), nil
}

func ClearSession(c *gin.Context) error {
	session := sessions.Default(c)
	session.Clear()
	return session.Save()
}
