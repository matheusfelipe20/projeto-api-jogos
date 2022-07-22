package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/router"
)

func main() {

	fmt.Println("Rodando")
	r := router.GerarRota()

	log.Fatal(http.ListenAndServe(":8080", r))

}
