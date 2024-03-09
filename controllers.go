package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lancatlin/lazy-finance/ledger"
	"github.com/lancatlin/lazy-finance/model"
)

// @Summary      Get Templates
// @Description  get templates for a user
// @Tags         templates
// @Accept       json
// @Produce      json
// @Success      200  {array}  string "Returns user templates"
// @Failure      500  {object}  Error  "Internal Server Error"
// @Router       /templates [get]
func getTemplates(c *gin.Context) {
	user := getUser(c)
	templates, err := user.templates()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, templates)
}

// @Summary      New Transaction
// @Description  create a new transaction
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        data  body      model.Transaction  true  "Transaction Data"
// @Success      200  {object}  string "Returns new transaction"
// @Failure      400  {object}  Error  "Bad Request"
// @Failure      500  {object}  Error  "Internal Server Error"
// @Router       /txs [post]
func newTx(c *gin.Context) {
	var tx model.Transaction
	if err := c.ShouldBind(&tx); err != nil {
		c.AbortWithError(400, err)
		return
	}
	user := getUser(c)
	err := user.newTx(tx)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.JSON(200, tx)
}

// @Summary      Get Transactions
// @Description  get transactions for a user
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        query  query      ledger.Query  false  "Query"
// @Success      200  {array}  model.Transaction  "Returns user transactions"
// @Failure      500  {object}  Error  "Internal Server Error"
// @Router       /txs [get]
func getTxs(c *gin.Context) {
	user := getUser(c)
	query := ledger.Query{}
	if err := c.Bind(&query); err != nil {
		c.AbortWithError(400, err)
		return
	}
	txs, err := user.txs(query)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, txs)
}

// @Summary      Get Balances
// @Description  get balances for a user
// @Tags         balances
// @Accept       json
// @Produce      json
// @Param        query  query      ledger.Query  false  "Query"
// @Success      200  {array}  ledger.Balance  "Returns user balances"
// @Failure      500  {object}  Error  "Internal Server Error"
// @Router       /balances [get]
func getBalances(c *gin.Context) {
	user := getUser(c)
	query := ledger.Query{}
	if err := c.Bind(&query); err != nil {
		c.AbortWithError(400, err)
		return
	}
	balances, err := user.getBalances(query)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, balances)
}

// @Summary      Get File List
// @Description  get file list for a user
// @Tags         files
// @Accept       json
// @Produce      json
// @Success      200  {array}  File  "Returns user file list"
// @Failure      500  {object}  Error  "Internal Server Error"
// @Router       /files [get]
func getFileList(c *gin.Context) {
	user := getUser(c)
	files, err := user.ListFiles()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, files)
}

// @Summary      Get File
// @Description  get file for a user
// @Tags         files
// @Accept       json
// @Produce      json
// @Param        path  path  string  true  "File Path"
// @Success      200  {file}  string  "Returns user file"
// @Failure      500  {object}  Error  "Internal Server Error"
// @Router       /files/{path} [get]
func getFile(c *gin.Context) {
	user := getUser(c)
	path := c.Param("path")
	c.File(user.FilePath(path))
}

// @Summary      Upload File
// @Description  upload file for a user
// @Tags         files
// @Accept       json
// @Produce      json
// @Param        path  path  string  true  "File Path"
// @Param        data  body  string  true  "File Data"
// @Success      200  {object}  string  "Returns success"
// @Failure      500  {object}  Error  "Internal Server Error"
// @Router       /files/{path} [post]
func uploadFile(c *gin.Context) {
	user := getUser(c)
	path := c.Param("path")
	var payload map[string]string
	if err := c.BindJSON(&payload); err != nil {
		c.AbortWithError(400, err)
		return
	}

	data, ok := payload["data"]
	if !ok {
		c.AbortWithError(400, fmt.Errorf("missing data"))
		return
	}
	err := user.overwriteFile(path, data)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}
