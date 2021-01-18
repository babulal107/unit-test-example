package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"
	appConfig "github.com/unit-test-example/internal/config"
	"github.com/unit-test-example/pkg/api/model"
	"github.com/unit-test-example/pkg/middleware"
	"github.com/unit-test-example/pkg/utils"
	"net/http"
)

// Init services, repositories, gin middleware and router
func Init(configs *appConfig.Application) *gin.Engine {

	r := gin.New()
	r.HandleMethodNotAllowed = utils.HandleMethodNotAllowed
	r.UseRawPath = utils.UseRawPath
	r.Use(gin.Recovery(), gin.Logger()) // default gin recovery and logger middleware used
	r.Use(func(context *gin.Context) {
		log.Debugf("%v", context.Request.Header)
		log.Debugf("Service Name %s", context.GetHeader(utils.HttpHeaderServiceName))
		context.Next()
	})
	r.Use(middleware.WithCacheHeaderControl("300")) // 5 mines

	// server status check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.Response{
			Code:    http.StatusOK,
			Message: "SERVICE UP",
			Data:    make([]interface{}, 0),
		})
	})

	return r
}
