package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeletarMissao(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"nome": "Missão 1",
	})
}
