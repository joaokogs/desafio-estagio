package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		//func(ctx *gin.Context) como se fosse a função anonima do js
		//Info da missao
		v1.GET("/missao", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"nome": "Missão 1",
			})
		})

		//Cria missões
		v1.POST("/missao", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"nome": "Missão 1",
			})
		})

		//Deleta uma missão
		v1.DELETE("/missao", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"nome": "Missão 1",
			})
		})

		//Edita uma missão
		v1.PUT("/missao", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"nome": "Missão 1",
			})
		})

		//Lista todas as missões
		v1.GET("/missoes", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"lista": "Missões",
			})
		})

	}
}
