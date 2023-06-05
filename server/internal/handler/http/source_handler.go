package handler

import (
	"codeview/config"
	"codeview/internal/dto/request"
	"codeview/internal/middleware"
	"codeview/internal/service"
	"codeview/internal/util/exception"
	"codeview/internal/util/pagination"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SourceHandler struct {
	cfg           config.AppConfig
	sourceService service.SourceService
}

func InitSourceHandler(cfg config.AppConfig, router *gin.Engine, sourceService service.SourceService) {
	h := &SourceHandler{
		cfg,
		sourceService,
	}

	g := router.Group("/api/source")
	g.Use(middleware.SessionAuth(cfg))
	g.POST("/", h.CreateSource)
	g.GET("/", h.GetSources)
}

func (h *SourceHandler) CreateSource(c *gin.Context) {
	var body request.SourceCreate

	err := c.ShouldBindJSON(&body)
	if err != nil {
		e := exception.NewBadRequest(fmt.Sprintf("Invalid request body: %s", err.Error()))
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	ctx := c.Request.Context()
	res, err := h.sourceService.CreateSource(ctx, &body)
	if err != nil {
		e := exception.NewBadRequest(err.Error())
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    res,
		"message": "success",
	})
}

func (h *SourceHandler) GetSources(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		e := exception.NewBadRequest("invalid page request params")
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil {
		e := exception.NewBadRequest("invalid size request params")
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	p := pagination.Parse(page, size)
	p.ValidatePagination()

	ctx := c.Request.Context()
	res, err := h.sourceService.GetSources(ctx, p)
	if err != nil {
		e := exception.NewBadRequest(err.Error())
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":        res,
		"currentPage": p.Page,
		"totalPage":   p.TotalPage,
		"pageSize":    p.PageSize,
		"totalSize":   p.Total,
		"message":     "success",
	})
}
