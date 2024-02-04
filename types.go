package main

type Error struct {
	Message string `json:"message"`
}

type TxData struct {
	Action      string `form:"action" binding:"required"`
	Name        string `form:"name"`
	Date        string
	Amount      string `form:"amount" binding:"required"`
	Destination string `form:"dest"`
	Source      string `form:"src"`
}

type TxResult struct {
	Tx string `json:"tx"`
}

type Query struct {
	Name  string `json:"name" binding:"required"`
	Query string `json:"query" binding:"required"`
}