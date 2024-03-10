package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/lancatlin/lazy-finance/docs"
)

func router() *gin.Engine {
	r := gin.Default()

	// Error handling middleware
	r.Use(errorHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")

	api.POST("/signup", signup)
	api.POST("/signin", signin)
	api.POST("/logout", logout)

	authApi := api.Group("", authenticate)

	authApi.GET("/templates", getTemplates)
	authApi.POST("/txs", newTx)
	authApi.GET("/txs", getTxs)
	authApi.GET("/balances", getBalances)
	authApi.GET("/files", getFileList)
	authApi.GET("/files/*path", getFile)
	authApi.POST("/files/*path", uploadFile)

	authZone := r.Group("", authenticate)

	authZone.GET("/download", func(c *gin.Context) {
		user := getUser(c)
		c.FileAttachment(user.FilePath(DEFAULT_JOURNAL), DEFAULT_JOURNAL)
	})

	return r
}
