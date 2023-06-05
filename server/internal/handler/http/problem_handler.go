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
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProblemHandler struct {
	cfg             config.AppConfig
	problemService  service.ProblemService
	questionService service.QuestionService
	solutionService service.SolutionService
}

func InitProblemHandler(
	cfg config.AppConfig,
	router *gin.Engine,
	problemService service.ProblemService,
	questionService service.QuestionService,
	solutionService service.SolutionService) {
	h := &ProblemHandler{
		cfg,
		problemService,
		questionService,
		solutionService,
	}

	g := router.Group("/api/problem")
	g.Use(middleware.SessionAuth(cfg))
	g.POST("/", h.CreateProblem)
	g.GET("/", h.GetProblems)
	g.GET("/list/", h.GetProblemsByIds)
	g.GET("/:id", h.GetProblemById)
	g.PATCH("/:id", h.UpdateProblemById)
	g.DELETE("/:id", h.DeleteProblemById)

	g_question := g.Group("/question")
	g_question.PATCH("/:id", h.UpdateProblemQuestionById)

	g_solution := g.Group("/solution")
	g_solution.POST("/", h.CreateProblemSolution)
	g_solution.PATCH("/:id", h.UpdateProblemSolutionById)
	g_solution.DELETE("/:id", h.DeleteProblemSolutionById)

}

func (h *ProblemHandler) CreateProblem(c *gin.Context) {
	var body request.ProblemCreate

	err := c.ShouldBindJSON(&body)
	if err != nil {
		e := exception.NewBadRequest(fmt.Sprintf("invalid request body: %s", err.Error()))
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
		"data": gin.H{
			"id": res.ID,
		},
		"message": "success",
	})
}

func (h *ProblemHandler) GetProblems(c *gin.Context) {
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

func (h *ProblemHandler) GetProblemById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		e := exception.NewBadRequest("invalid ID params")
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

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

func (h *ProblemHandler) GetProblemsByIds(c *gin.Context) {
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

	var ids []uuid.UUID

	for _, s := range strings.Split(c.Query("ids"), ",") {
		id, err := uuid.Parse(s)
		if err != nil {
			e := exception.NewBadRequest(fmt.Sprintf("invalid ID %s: %v", s, err))
			c.JSON(e.Status(), gin.H{"error": e})
			return
		}
		ids = append(ids, id)
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

func (h *ProblemHandler) UpdateProblemById(c *gin.Context) {
	var body request.ProblemUpdate

	err := c.ShouldBindJSON(&body)
	if err != nil {
		e := exception.NewBadRequest(fmt.Sprintf("invalid request body: %s", err.Error()))
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		e := exception.NewBadRequest("invalid ID params")
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

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

func (h *ProblemHandler) DeleteProblemById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		e := exception.NewBadRequest("invalid ID params")
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	ctx := c.Request.Context()
	err = h.problemService.DeleteProblemById(ctx, id)
	if err != nil {
		e := exception.NewBadRequest(err.Error())
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "success",
	})
}

func (h *ProblemHandler) UpdateProblemQuestionById(c *gin.Context) {
	var body request.QuestionUpdate

	err := c.ShouldBindJSON(&body)
	if err != nil {
		e := exception.NewBadRequest(fmt.Sprintf("invalid request body: %s", err.Error()))
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		e := exception.NewBadRequest("invalid ID params")
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	ctx := c.Request.Context()
	_, err = h.questionService.UpdateQuestionById(ctx, id, &body)
	if err != nil {
		e := exception.NewBadRequest(err.Error())
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *ProblemHandler) CreateProblemSolution(c *gin.Context) {
	var body request.SolutionCreate

	err := c.ShouldBindJSON(&body)
	if err != nil {
		e := exception.NewBadRequest(fmt.Sprintf("invalid request body: %s", err.Error()))
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	ctx := c.Request.Context()
	_, err = h.solutionService.CreateSolution(ctx, &body)
	if err != nil {
		e := exception.NewBadRequest(err.Error())
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *ProblemHandler) UpdateProblemSolutionById(c *gin.Context) {
	var body request.SolutionUpdate

	err := c.ShouldBindJSON(&body)
	if err != nil {
		e := exception.NewBadRequest(fmt.Sprintf("invalid request body: %s", err.Error()))
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		e := exception.NewBadRequest("invalid ID params")
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	ctx := c.Request.Context()
	_, err = h.solutionService.UpdateSolutionById(ctx, id, &body)
	if err != nil {
		e := exception.NewBadRequest(err.Error())
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *ProblemHandler) DeleteProblemSolutionById(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		e := exception.NewBadRequest("invalid ID params")
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	ctx := c.Request.Context()
	err = h.solutionService.DeleteSolutionById(ctx, id)
	if err != nil {
		e := exception.NewBadRequest(err.Error())
		c.JSON(e.Status(), gin.H{"error": e})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "success",
	})
}
