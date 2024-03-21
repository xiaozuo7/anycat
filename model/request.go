package model

type TransReq struct {
	Content string `json:"content" binding:"required"`
	From    string `json:"from" binding:"required" default:"auto"`
	To      string `json:"to" binding:"required"`
}
