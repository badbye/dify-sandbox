package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/langgenius/dify-sandbox/internal/middleware"
	"github.com/langgenius/dify-sandbox/internal/static"
)

func Setup(eng *gin.Engine) {
	eng.Use(middleware.MaxRequest(static.GetCoshubGlobalConfigurations().MaxRequests))
	eng.Use(middleware.Auth())
	eng.Use(middleware.MaxWorker(static.GetCoshubGlobalConfigurations().MaxWorkers))

	eng.POST("/v1/sandbox/run", RunSandboxController)
}
