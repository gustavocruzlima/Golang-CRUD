package main

import (
	"net/http"

	"github.com/gustavocruzlima/web/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
