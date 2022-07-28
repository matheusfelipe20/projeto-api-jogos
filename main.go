package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matheusfelipe20/projeto-api-jogos/src/configs"
	"github.com/matheusfelipe20/projeto-api-jogos/src/router"
)

func main() {

	configs.Carregar()
	r := router.GerarRota()

	fmt.Printf("Escutando na porta %d\n", configs.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configs.Porta), r))

}
