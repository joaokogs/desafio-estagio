package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaokogs/desafio-estagio/handler"
)

func InitRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		//Info da missao
		v1.GET("/missao", handler.MostrarMissao)

		//Cria missões
		v1.POST("/missao", handler.CriarMissao)

		//Deleta uma missão
		v1.DELETE("/missao", handler.DeletarMissao)

		//Edita uma missão
		v1.PUT("/missao", handler.AtualizarMissao)

		//Lista todas as missões
		v1.GET("/missoes", handler.ListarMissoes)

	}
}
