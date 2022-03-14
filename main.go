package main

import (
	"github.com/contecloud/gin-api-rest/database"
	"github.com/contecloud/gin-api-rest/models"
	"github.com/contecloud/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Carlo Conte", CPF: "01199294063", RG: "5085403946"},
		{Nome: "Ana", CPF: "01312312323", RG: "505216"},
		{Nome: "Sérgio", CPF: "32132132121", RG: "123465479"},
	}
	routes.HandleRequests()
}
