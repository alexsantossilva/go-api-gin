package main

import (
	"github.com/alexsantossilva/go-api-gin/database"
	//"github.com/alexsantossilva/api-go-gin/models"
	"github.com/alexsantossilva/go-api-gin/routes"
)

func main() {
	database.ConnectDB()
	//models.Alunos = []models.Aluno{
	//	{Nome: "Alex", RG: "23423432", CPF: "1234567809"},
	//	{Nome: "Maria", RG: "23423482", CPF: "1234567812"},
	//	{Nome: "Jose", RG: "23423437", CPF: "12345678100"},
	//	{Nome: "Rola", RG: "23423436", CPF: "1234567822"},
	//}
	routes.HandleRequests()
}
