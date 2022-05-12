package main

import (
	docs "demo/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	r.Run(":3000")
}

// @BasePath /api/v1
// @Tags Swagger Demo
// @Summary ping demo
// @Description handle ping request
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Router /ping [get]
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

// @Summary echo demo
// @Description echo input request
// @Tags Swagger Demo
// @Param name body []main.User true "User information"
// @Accept json
// @Produce json
// @Success 200 {array} main.User "OK"
// @Router /echo [post]
func Echo(c *gin.Context) {
	/*
		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusNotFound, "Failed")
		}
		c.JSON(http.StatusOK, data)
	*/
	var u []User
	if c.Bind(&u) == nil {
		c.JSON(http.StatusOK, u)
	} else {
		c.JSON(http.StatusBadRequest, "Failed")
	}
}
