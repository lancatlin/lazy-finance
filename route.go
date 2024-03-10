package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/lancatlin/lazy-finance/docs"
)

func router() *gin.Engine {
	r := gin.Default()

	r.Use(errorHandler)
	r.Use(session)

	r.NoRoute(staticFiles)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")

	api.POST("/signup", signup)
	api.POST("/signin", signin)
	api.POST("/logout", logout)
	api.GET("/status", status)

	authApi := api.Group("", authRequired)

	authApi.GET("/templates", getTemplates)
	authApi.POST("/txs", newTx)
	authApi.GET("/txs", getTxs)
	authApi.GET("/balances", getBalances)
	authApi.GET("/files", getFileList)
	authApi.GET("/files/*path", getFile)
	authApi.POST("/files/*path", uploadFile)

	authZone := r.Group("", session)

	authZone.GET("/download", func(c *gin.Context) {
		user := getUser(c)
		c.FileAttachment(user.FilePath(DEFAULT_JOURNAL), DEFAULT_JOURNAL)
	})

	return r
}
