package main

import (
	"github/DarleneCosta/api-go-gin/database"
	"github/DarleneCosta/api-go-gin/models"
	"github/DarleneCosta/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Marcelo Lima", CPF: "00000000000", RG: "450000523"},
		{Nome: "Roberta Souza", CPF: "11111111111", RG: "470000692"},
	}
	routes.HandleRequests()
}
