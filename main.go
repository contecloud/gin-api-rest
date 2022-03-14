package main

import (
	"github.com/contecloud/gin-api-rest/database"
	"github.com/contecloud/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
