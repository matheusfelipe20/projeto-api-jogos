package validators

import (
	"log"
	"time"
)

//Validador para limitar a venda do jogo a 1 dia antes do evento
func ValidadeVenda(data string) bool {

	dataJogo := data

	parsed, err := time.Parse("02-01-2006", dataJogo)
	if err != nil {
		log.Fatal(err)
	}
	beforeDay := parsed.AddDate(0, 0, 1) //Somar 1 dia a data do evento
	today := time.Now()

	comparetedEvento := beforeDay.Before(today)

	return comparetedEvento
}
