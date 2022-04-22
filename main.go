package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io/ioutil"
	"net/http"
	docs "wayfeng/gin-swag-demo/docs"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	api_v1 := r.Group("/api/v1")
	api_v1.GET("/ping", Pong)
	api_v1.POST("/echo", Echo)

	r.HEAD("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}

// @BasePath /api/v1
// PingDemo godoc
// @Summary ping demo
// @Schemes
// @Description handle ping request
// @Tags demo
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Router /ping [get]
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

// @BasePath /api/v1
// PingDemo godoc
// @Summary echo demo
// @Schemes
// @Description echo input request
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Tags demo
// @Accept json
// @Produce json
// @Success 200 {string} json {}
// @Router /echo [post]
func Echo(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusNotFound, "Failed")
	}
	c.JSON(http.StatusOK, data)
}
