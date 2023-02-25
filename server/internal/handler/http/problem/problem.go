package problemhandler

import (
	"codeview/internal/dto/request"
	"codeview/internal/exception"
	"codeview/utils/pagination"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateProblem(c *gin.Context) {
	var body request.ProblemCreate

	err := c.ShouldBindJSON(&body)
	if err != nil {
		e := exception.NewBadRequest(fmt.Sprintf("Invalid request body: %s", err.Error()))
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	ctx := c.Request.Context()
	res, err := h.problemService.CreateProblem(ctx, &body)
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

func (h *Handler) GetProblems(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		e := exception.NewBadRequest("Invalid page request params")
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil {
		e := exception.NewBadRequest("Invalid size request params")
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	p := pagination.Parse(page, size)
	p.ValidatePagination()

	ctx := c.Request.Context()
	res, err := h.problemService.GetProblems(ctx, p)
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

func (h *Handler) GetProblemById(c *gin.Context) {
	u64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		e := exception.NewBadRequest("Invalid ID params")
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	id := uint(u64)

	ctx := c.Request.Context()
	res, err := h.problemService.GetProblemById(ctx, id)
	if err != nil {
		e := exception.NewBadRequest(err.Error())
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "success",
	})
}

func (h *Handler) GetProblemsByIds(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		e := exception.NewBadRequest("Invalid page request params")
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil {
		e := exception.NewBadRequest("Invalid size request params")
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	p := pagination.Parse(page, size)
	p.ValidatePagination()

	var ids []uint

	for _, s := range strings.Split(c.Query("ids"), ",") {
		u64, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			e := exception.NewBadRequest("Invalid ID")
			c.JSON(e.Status(), gin.H{"error": e})
			return
		}
		ids = append(ids, uint(u64))
	}

	ctx := c.Request.Context()
	res, err := h.problemService.GetProblemsByIds(ctx, ids, p)
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

func (h *Handler) UpdateProblemById(c *gin.Context) {
	var body request.ProblemUpdate

	err := c.ShouldBindJSON(&body)
	if err != nil {
		e := exception.NewBadRequest(fmt.Sprintf("Invalid request body: %s", err.Error()))
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	u64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		e := exception.NewBadRequest("Invalid ID params")
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	id := uint(u64)

	ctx := c.Request.Context()
	_, err = h.problemService.UpdateProblemById(ctx, id, &body)
	if err != nil {
		e := exception.NewBadRequest(err.Error())
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}

func (h *Handler) DeleteProblemById(c *gin.Context) {
	u64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		e := exception.NewBadRequest("Invalid ID params")
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	id := uint(u64)

	ctx := c.Request.Context()
	err = h.problemService.DeleteProblemById(ctx, id)
	if err != nil {
		e := exception.NewBadRequest(err.Error())
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}
