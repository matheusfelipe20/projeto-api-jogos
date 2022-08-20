package validators

import "testing"

// TestVerificarCPF verifica se um cpf é válido
func TestVerificarCPF(t *testing.T) {

	cpf1 := "984.618.390-99" // String normal
	cpf2 := "18255802002"    // String com digito verificador 0
	cpf3 := 21240881061      // integer normal
	cpf4 := 3501998902       // integer faltando digito - inválido
	cpf5 := 43527894020      // integer com verificador 0
	cpf6 := "4201767000"     // String faltando digito - inválido

	if verificar1, _ := VerificarCPFbyString(cpf1); verificar1 != true {
		t.Errorf("CPF %s não é válido", cpf1)
	}

	if verificar2, _ := VerificarCPFbyString(cpf2); verificar2 != true {
		t.Errorf("CPF %s não é válido", cpf2)
	}

	if verificar3 := VerificarCPF(cpf3); verificar3 != true {
		t.Errorf("CPF %d não é válido", cpf3)
	}

	if verificar4 := VerificarCPF(cpf4); verificar4 != false {
		t.Errorf("CPF %d não pode ser válido", cpf4)
	}

	if verificar5 := VerificarCPF(cpf5); verificar5 != true {
		t.Errorf("CPF %d não é válido", cpf5)
	}

	if verificar6, _ := VerificarCPFbyString(cpf6); verificar6 != false {
		t.Errorf("CPF %s não pode ser válido", cpf6)
	}
}
