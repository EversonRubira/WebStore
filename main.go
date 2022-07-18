package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=Alura_Webstore password=Postgres202201* host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Nome       string
	Descriçao  string
	Preço      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("view/templates/*.html"))

func main() {
	db := conectaComBancoDeDados()
	defer db.Close()
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "camiseta", Descriçao: "Especial", Preço: 39, Quantidade: 10},
		{"tenis", "Corrida", 89, 3},
		{"Fone", "Sem Fios", 99, 5},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
