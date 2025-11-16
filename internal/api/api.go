package api

import (
	"log"

	"net/http"

	"github.com/developerasun/SignalDash/pkg"
	"github.com/gin-gonic/gin"
)

// Health godoc
// @Summary Show the health status
// @Description Get server health status
// @Tags api
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /api/health [get]
func Health(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, HealthResponse{
		Message: "ok",
	})
}

// CrawlDollarIndex godoc
// @Summary visit tradingview and extract daily dollar index
// @Description Get server health status
// @Tags api
// @Produce json
// @Success 200 {object} CrawlResponse
// @Router /api/dxy [get]
func CrawlDollarIndex(ctx *gin.Context) {
	dxy, rcErr := pkg.CrawlDollarIndex()

	if rcErr != nil {
		log.Println(`CrawlDollarIndex: failed to run a crawler`)
		ctx.Error(rcErr)
	}

	ctx.JSON(http.StatusOK, CrawlResponse{
		Data: dxy,
	})
}

// CrawlDaiPrice godoc
// @Summary visit metamask and extract dai token price at the moment
// @Description Get dai coin price
// @Tags api
// @Produce json
// @Success 200 {object} CrawlResponse
// @Router /api/dai [get]
func CrawlDaiPrice(ctx *gin.Context) {
	dai, rcErr := pkg.CrawlDaiPrice()

	if rcErr != nil {
		log.Println(`CrawlDaiPrice: failed to run a crawler`)
		ctx.Error(rcErr)
	}

	ctx.JSON(http.StatusOK, CrawlResponse{
		Data: dai,
	})
}

// RenderMainPage godoc
// @Summary show main page, returning html
// @Description show main page, returning html
// @Tags view
// @Router / [get]
func RenderMainPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
