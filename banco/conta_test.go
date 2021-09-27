package banco

import (
	"fmt"
	"sync"
	"testing"
)

func TestCriarConta(t *testing.T) {
	c := novaConta()

	impressaoContaObtida := fmt.Sprintf("%s", c)
	impressaoContaEsperada := "Conta: 0001 de 'Paulo', SALDO: R$ 100.00"

	if impressaoContaObtida != impressaoContaEsperada {
		t.Errorf("Impressão de conta esperada diferente da obtida. Esperado: '%s', Obtido: '%s'", impressaoContaEsperada, impressaoContaObtida)
	}

	impressaoContaObtida = c.String()

	if impressaoContaObtida != impressaoContaEsperada {
		t.Errorf("Impressão de conta obtida diferente da esperada. Esperado: '%s', Obtido: '%s'", impressaoContaEsperada, impressaoContaObtida)
	}
}

func TestDeposito(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	c := novaConta()
	saldoOriginal := c.saldo

	valorDeposito := Monetario(200)
	go c.Deposito(valorDeposito, &wg)
	wg.Wait()

	saldoEsperado := saldoOriginal + valorDeposito
	saldoObtido := c.saldo
	if saldoObtido != saldoEsperado {
		t.Errorf("Saldo obtido diferente do esperado. Esperado: %s, Obtido: %s", saldoEsperado, saldoObtido)
	}
}

func TestSaque(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	c := novaConta()
	saldoOriginal := c.saldo

	valorSaque := Monetario(75)
	go c.Saque(valorSaque, &wg)
	wg.Wait()

	saldoEsperado := saldoOriginal - valorSaque
	saldoObtido := c.saldo
	if saldoObtido != saldoEsperado {
		t.Errorf("Saldo obtido diferente do esperado. Esperado: %s, Obtido: %s", saldoEsperado, saldoObtido)
	}
}

func TestSaldo(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	c := novaConta()
	saldoOriginal := c.saldo

	valorDeposito := Monetario(200)
	valorSaque := Monetario(75)

	go c.Deposito(valorDeposito, &wg)
	go c.Saque(valorSaque, &wg)

	wg.Wait()

	saldoEsperado := saldoOriginal + valorDeposito - valorSaque
	saldoObtido := c.saldo
	if saldoObtido != saldoEsperado {
		t.Errorf("Saldo obtido diferente do esperado. Esperado: %s, Obtido: %s", saldoEsperado, saldoObtido)
	}
}

func novaConta() *Conta {
	return NewConta("Paulo", "0001", 100)
}
