package handler

import (
	"fmt"
	"net/http"

	"github.com/AykoSousa/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Operation from handler: %s successfull", operation),
		"data": data,
	})
}

type ErrorResponse struct {
	Message string `json: "message"`
	ErrorCode string `json: "errorCode"`
}

type CreateOpeningResponse struct {
	Message string `json: "message"`
	Data schemas.OpeningResponse `json: "data"`
}

type DeleteOpeningResponse struct {
	Message string `json: "message"`
	Data schemas.OpeningResponse `json: "data"`
}