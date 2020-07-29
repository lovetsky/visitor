package account

import (
	"fmt"

	"github.com/labstack/gommon/color"
)

type visiter interface {
	VisitFoAccount(int int)
}

type accounter interface {
	Reserved(uidAccount string) (status bool)
	Accept(v visiter)
}

type account struct {
	uid     string
	balance int
	visiter visiter
}

// Reserved реализует метод резервирования товара покупателем
func (a *account) Reserved(uidAccount string) (status bool) {
	done := color.Green("[ OK ]")
	fmt.Println(done, " Товар был зарезервирован для покупателя: ", uidAccount)
	return true
}

// Accept осуществляет вызов новой функции для типа
func (a *account) Accept(v visiter) {
	a.visiter.VisitFoAccount(1)
}

// NewAccount создает реализацию интерефейса accounter
func NewAccount(uidAccount string, visiter visiter) accounter {
	return &account{
		uid:     uidAccount,
		balance: 100,
		visiter: visiter,
	}
}
