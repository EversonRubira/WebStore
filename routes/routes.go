package routes

import (
	"net/http"
)

func carregaRotas() {

	http.HandleFunc("/", controllers.Index)
}
