package banco

import (
	"fmt"
	"sync"
)

type Monetario float64

func (m Monetario) String() string {
	return fmt.Sprintf(
		"R$ %.2f",
		m,
	)
}

type Conta struct {
	nome   string
	numero string
	saldo  Monetario
	mutex  sync.Mutex
}

func (c *Conta) String() string {
	return fmt.Sprintf(
		"Conta: %s de '%s', SALDO: %s",
		c.numero, c.nome, c.saldo)
}

func NewConta(nome, numero string, saldo float64) *Conta {
	return &Conta{
		nome:   nome,
		numero: numero,
		saldo:  Monetario(saldo),
	}
}

func (c *Conta) Deposito(valor Monetario, wg *sync.WaitGroup) {
	c.mutex.Lock()
	c.saldo += valor
	c.mutex.Unlock()
	wg.Done()
}

func (c *Conta) Saque(valor Monetario, wg *sync.WaitGroup) {
	c.mutex.Lock()
	c.saldo -= valor
	c.mutex.Unlock()
	wg.Done()
}
