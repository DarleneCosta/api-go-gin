package main

import (
	"github/DarleneCosta/api-go-gin/database"
	"github/DarleneCosta/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
