package validators

import (
	"log"
	"strconv"
	"time"
)

//Validador para verificar se o usuario é maior de idade
func ValidadeDataNascimento(nascimento string) bool {

	dataUser := nascimento

	parsed, err := time.Parse("02-01-2006", dataUser)
	if err != nil {
		log.Fatal(err)
	}
	beforeYear := parsed.AddDate(18, 0, 0) //Somar 18 anos a data de nascimento do usuario
	today := time.Now()

	compareted := beforeYear.Before(today)

	return compareted
}

func ValidadeCPF(cpf int) bool {
	cpfUser := cpf

	quantityCPF := strconv.Itoa(cpfUser)

	//o cpf será verificado se tem 11 digitos
	return len(quantityCPF) == 11
}

func ValidadeNome(nome string) bool {
	nomeUser := nome

	//o nome será verificado para saaber se é vazio ou não
	return nomeUser != ""
}
