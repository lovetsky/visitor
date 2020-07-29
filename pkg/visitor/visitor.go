package visitor

import (
	"fmt"
)

type VisitorAccepter interface {
	Accept(v Visitor) (err error)
}

// Visitor обеспечивает интерфейс покупателя
type Visitor interface {
	VisitFoAccount(l VisitorAccepter) (err error)
}

type calcAddons struct {
	money int
}

// VisitFoAccount добавляем новый метод
func (a *calcAddons) VisitFoAccount(l VisitorAccepter) (err error) {
	fmt.Println("Добавили денег на кошелек (новый метод):", a.money)
	return
}

// NewVisitor фабрика визитора
func NewVisitor(money int) Visitor {
	return &calcAddons{
		money,
	}
}
