package handler

import (
	"codeview/domain"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	problemService domain.ProblemService
	config         *Config
}

type Config struct {
}

func New(router *gin.Engine, problemService domain.ProblemService, config *Config) {
	h := &Handler{
		problemService,
		config,
	}

	g := router.Group("/api/problem")
	g.POST("/", h.CreateProblem)
	g.GET("/", h.GetProblems)
	g.GET("/list/", h.GetProblemsByIds)
	g.GET("/:id", h.GetProblemById)
	g.PATCH("/:id", h.UpdateProblemById)
	g.DELETE("/:id", h.DeleteProblemById)

}
