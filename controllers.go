package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lancatlin/lazy-finance/model"
)

// @Summary      Get Queries
// @Description  get queries for a user
// @Tags         queries
// @Accept       json
// @Produce      json
// @Success      200  {array}  Query  "Returns user queries"
// @Failure      500  {object}  Error  "Internal Server Error"
// @Router       /queries [get]
func getQueries(c *gin.Context) {
	user := getUser(c)
	queries, err := user.queries()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, queries)
}

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
