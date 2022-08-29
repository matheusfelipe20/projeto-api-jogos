package models

import (
	"errors"

	"github.com/matheusfelipe20/projeto-api-jogos/src/models/validators"
)

type Vendas struct {
	Id                 int     `json:"id,omitempty"`
	Id_jogo            int     `json:"id_jogo,omitempty"`
	Titulo_jogo        string  `json:"titulo_jogo,omitempty"`
	Campeonato         string  `json:"campeonato,omitempty"`
	Data_jogo          string  `json:"data_jogo,omitempty"`
	Opcao_aposta       string  `json:"opcao_aposta,omitempty"`
	Valor_aposta       float64 `json:"valor_aposta,omitempty"`
	Limite_aposta      int     `json:"limite_aposta,omitempty"`	//validar limite de aposta 
	Cliente_nome       string  `json:"cliente_nome,omitempty"`
	Cliente_cpf        string  `json:"cliente_cpf,omitempty"`
	Cliente_nascimento string  `json:"cliente_nascimento,omitempty"`
}

func (vd *Vendas) ValidarVendas() error {

	if verificarIdJogo := validators.ValidadeID(uint64(vd.Id_jogo)); !verificarIdJogo {
		return errors.New("id do jogo é igual a 0")
	}
	if verificarTituloJogo := validators.ValidarCampo(vd.Titulo_jogo); !verificarTituloJogo {
		return errors.New("falha ao  cadastrar, insira o titulo do jogo")
	}
	if verificarCampeonato := validators.ValidarCampo(vd.Campeonato); !verificarCampeonato{
		return errors.New("falha ao cadastrar, insira o campeonato")
	}
	/* if verificarDataJogo := validators.ValidadeDataVenda(vd.Data_jogo); !verificarDataJogo {
		return errors.New("falha ao cadastrar, insira a data do jogo")
	} */
	if verificarNomeCliente := validators.ValidarCampo(vd.Cliente_nome); !verificarNomeCliente {
		return errors.New("falha ao cadastrar, insira o nome do cliente")
	}
	if verificarCpfCliente, _ := validators.VerificarCPFbyString(vd.Cliente_cpf); !verificarCpfCliente {
		return errors.New("falha ao cadastrar, cpf inválido")
	}
	/* if verificarNomeCliente := validators.ValidadeDataNascimento(vd.Cliente_nascimento); !verificarNomeCliente {
		return errors.New("falha ao cadastrar, usuário menor de idade")
	} */

	return nil
}
