package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"codeview/domain"
	"codeview/exception"
	"codeview/validation"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	problemService domain.ProblemService
	config         *Config
}

type Config struct {
	TimeoutDuration time.Duration
	MaxBodyBytes    int64
}

func New(router *gin.Engine, problemService domain.ProblemService, config *Config) {
	h := &Handler{
		problemService,
		config,
	}

	g := router.Group("/api/problem")
	g.POST("/image", h.Image)
	g.GET("/:id", h.GetProblemById)
}

func (h *Handler) GetProblemById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Problem ID", id)

	c.JSON(http.StatusOK, gin.H{
		"title":      "New Problem",
		"difficulty": gin.H{"label": "easy", "color": "teal"},
		"rating":     4,
		"tags":       []string{},
		"sources":    []string{},
		"question": gin.H{
			"content": `<h2 style="text-align: center; margin-left: 0px!important;"><strong>Welcome to Code View</strong></h2>`,
		},
		"solution": []string{},
	})

}

func (h *Handler) Image(c *gin.Context) {
	// limit overly large request bodies
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, h.config.MaxBodyBytes)

	imageFileHeader, err := c.FormFile("imageFile")

	// check for error before checking for non-nil header
	if err != nil {
		// should be a validation error
		log.Printf("Unable parse multipart/form-data: %+v", err)

		if err.Error() == "http: request body too large" {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{
				"error": fmt.Sprintf("Max request body size is %v bytes\n", h.config.MaxBodyBytes),
			})
			return
		}

		e := exception.NewBadRequest("Unable to parse multipart/form-data")
		c.JSON(e.Status(), gin.H{
			"error": e,
		})

		return
	}

	if imageFileHeader == nil {
		e := exception.NewBadRequest("Must include an imageFile")
		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	mimeType := imageFileHeader.Header.Get("Content-Type")
	if valid := validation.IsAllowedImageType(mimeType); !valid {
		log.Println("Image is not an allowable mime-type")
		e := exception.NewBadRequest("imageFile must be 'image/jpeg' or 'image/png'")
		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	ctx := c.Request.Context()
	image, err := h.problemService.AddProblemImage(ctx, imageFileHeader)
	if err != nil {
		c.JSON(exception.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"imageUrl": image.URL,
		"message":  "success",
	})
}
