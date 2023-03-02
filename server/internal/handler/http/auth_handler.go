package handler

import (
	"codeview/config"
	"codeview/internal/dto/request"
	"codeview/internal/service"
	"codeview/internal/util"
	"codeview/internal/util/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	cfg         config.AppConfig
	authService service.AuthService
}

func InitAuthHandler(cfg config.AppConfig, router *gin.Engine, authService service.AuthService) {

	h := &AuthHandler{
		cfg,
		authService,
	}

	g := router.Group("/api/auth")
	g.POST("/register", h.Register)
	g.POST("/login", h.Login)

}

func (h *AuthHandler) Register(c *gin.Context) {
	var body request.Register

	if err := c.ShouldBindJSON(&body); err != nil {
		e := exception.NewBadRequest(err.Error())
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	ctx := c.Request.Context()
	token, err := h.authService.Register(ctx, &body)
	if err != nil {
		e := exception.NewBadRequest(err.Error())
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	if err := util.SetTokenSession(c, token); err != nil {
		e := exception.NewInternal()
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var body request.Login

	if err := c.ShouldBindJSON(&body); err != nil {
		e := exception.NewBadRequest(err.Error())
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	ctx := c.Request.Context()
	token, err := h.authService.Login(ctx, &body)
	if err != nil {
		e := exception.NewBadRequest(err.Error())
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	if err := util.SetTokenSession(c, token); err != nil {
		e := exception.NewInternal()
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (handler *AuthHandler) Logout(c *gin.Context) {

	if err := util.ClearSession(c); err != nil {
		e := exception.NewInternal()
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
