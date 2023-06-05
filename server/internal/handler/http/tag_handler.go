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

type TagHandler struct {
	cfg        config.AppConfig
	tagService service.TagService
}

func InitTagHandler(cfg config.AppConfig, router *gin.Engine, tagService service.TagService) {
	h := &TagHandler{
		cfg,
		tagService,
	}

	g := router.Group("/api/tag")
	g.Use(middleware.SessionAuth(cfg))
	g.POST("/", h.CreateTag)
	g.GET("/", h.GetTags)
}

func (h *TagHandler) CreateTag(c *gin.Context) {
	var body request.TagCreate

	err := c.ShouldBindJSON(&body)
	if err != nil {
		e := exception.NewBadRequest(fmt.Sprintf("Invalid request body: %s", err.Error()))
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	ctx := c.Request.Context()
	res, err := h.tagService.CreateTag(ctx, &body)
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

func (h *TagHandler) GetTags(c *gin.Context) {
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

	tagType := c.DefaultQuery("type", "PROBLEM")

	p := pagination.Parse(page, size)
	p.ValidatePagination()

	ctx := c.Request.Context()
	res, err := h.tagService.GetTagsByType(ctx, tagType, p)
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
