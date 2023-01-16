package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/webtoon/internal/app/request"
	"github.com/webtoon/internal/app/service"
)

type ManhwaController struct {
	manhwaService service.ManhwaService
}

func NewController(manhwaService service.ManhwaService) *ManhwaController {
	return &ManhwaController{
		manhwaService: manhwaService,
	}
}

func (c *ManhwaController) SetupRoutes(router *gin.Engine) {
	router.GET("/manhwas/:title", c.GetManhwaByTitle)
	router.GET("/manhwas", c.GetManhwas)
	router.POST("/manhwas", c.CreateManhwa)
}

func (c *ManhwaController) GetManhwaByTitle(ctx *gin.Context) {
	manhwaTitle := ctx.Param("title")
	manhwa, err := c.manhwaService.GetManhwaByTitle(manhwaTitle)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if manhwa == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Manhwa not found"})
		return
	}
	ctx.JSON(http.StatusOK, manhwa)
}

func (c *ManhwaController) GetManhwas(ctx *gin.Context) {
	manhwas, err := c.manhwaService.GetManhwas()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, manhwas)
}

func (c *ManhwaController) CreateManhwa(ctx *gin.Context) {
	var request request.CreateManhwaRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	manhwa, err := c.manhwaService.CreateManhwa(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}
