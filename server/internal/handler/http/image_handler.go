package handler

import (
	"codeview/config"
	"codeview/internal/service"
	"codeview/internal/util"
	"codeview/internal/util/exception"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	cfg          config.AppConfig
	imageService service.ImageService
}

func InitImageHandler(cfg config.AppConfig, router *gin.Engine, imageService service.ImageService) {
	h := &ImageHandler{
		cfg,
		imageService,
	}

	g := router.Group("/api/image")
	g.POST("/upload", h.Image)
}

func (h *ImageHandler) Image(c *gin.Context) {
	// limit overly large request bodies
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, h.cfg.MaxBodyBytes)

	imageFileHeader, err := c.FormFile("imageFile")

	// check for error before checking for non-nil header
	if err != nil {
		// should be a validation error
		log.Printf("Unable parse multipart/form-data: %+v", err)

		if err.Error() == "http: request body too large" {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{
				"error": fmt.Sprintf("Max request body size is %v bytes\n", h.cfg.MaxBodyBytes),
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
	if valid := util.IsAllowedImageType(mimeType); !valid {
		log.Println("Image is not an allowable mime-type")
		e := exception.NewBadRequest("imageFile must be 'image/jpeg' or 'image/png'")
		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	ctx := c.Request.Context()
	image, err := h.imageService.UploadImage(ctx, imageFileHeader)
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
