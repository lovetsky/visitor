package visitor

import (
	"fmt"
	"github.com/lovetsky/visitor/pkg/account"
)

// Visitor обеспечивает интерфейс покупателя
type Visitor interface {
	VisitFoAccount(l account.Accounter) (err error)
}

type calcAddons struct {
	money int
}

// VisitFoAccount добавляем новый метод
func (a *calcAddons) VisitFoAccount(l account.Accounter) (err error) {
	fmt.Println("Добавили денег на кошелек (новый метод):", a.money)
	return
}

// NewVisitor фабрика визитора
func NewVisitor(money int) Visitor {
	return &calcAddons{
		money,
	}
}
