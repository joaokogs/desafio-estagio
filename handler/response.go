package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaokogs/desafio-estagio/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type","aplication/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})

}

func sendSuccess(ctx *gin.Context, op string,data interface{}){
	ctx.Header("Content-type","aplication/json")
	ctx.JSON(http.StatusOK, gin.H{
		
		"message": fmt.Sprintf("operation from handler %s successfull", op),
		"data":data,
		
	})
}

type ErrorResponse struct{
	Message string `json: "message"`
	ErrorCode string `json : "errorCode"`

}

type CreateMissionResponse struct{
	Message string `json : "message"`
	Data schemas.MissionsOpening `json: "data"`
}

type DeleteMissionResponse struct {
	Message string                  `json:"message"`
	Data    schemas.MissionsOpening `json:"data"`
}
type ShowMissionResponse struct {
	Message string                  `json:"message"`
	Data    schemas.MissionsOpening `json:"data"`
}
type ListMissionsResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.MissionsOpening `json:"data"`
}
type UpdateMissionResponse struct {
	Message string                  `json:"message"`
	Data    schemas.MissionsOpening `json:"data"`
}