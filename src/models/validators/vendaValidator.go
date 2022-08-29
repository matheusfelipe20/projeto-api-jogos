package validators

import (
	"log"
	"time"
)


//Validador para limitar a venda do jogo a 1 dia antes do evento
func ValidadeDataVenda(data string) bool {


	if data != "" {
		return false
	}

	parsed, err := time.Parse("02-01-2006", data)
	if err != nil {
		log.Println(err)
		return false
	}
	beforeDay := parsed.AddDate(0, 0, 1) //Somar 1 dia a data do evento
	today := time.Now()

	comparetedEvento := beforeDay.Before(today)

	return comparetedEvento
}

