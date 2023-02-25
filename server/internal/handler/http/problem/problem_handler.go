package problemhandler

import (
	"codeview/config"
	"codeview/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	config         config.AppConfig
	problemService service.ProblemService
}

func New(config config.AppConfig, router *gin.Engine, problemService service.ProblemService) {
	h := &Handler{
		config,
		problemService,
	}

	g := router.Group("/api/problem")
	g.POST("/", h.CreateProblem)
	g.GET("/", h.GetProblems)
	g.GET("/list/", h.GetProblemsByIds)
	g.GET("/:id", h.GetProblemById)
	g.PATCH("/:id", h.UpdateProblemById)
	g.DELETE("/:id", h.DeleteProblemById)

}
