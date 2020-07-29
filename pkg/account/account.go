package account

import (
	"fmt"
	"github.com/lovetsky/visitor/pkg/visitor"
)

// Accounter интерфейс покупателя
type Accounter interface {
	Check(uidAccount string) (status bool)
	Accept(v visitor.Visitor) (err error)
}

type account struct {
	uid     string
	balance int
}

// Check реализует метод проверки пользователя
func (a *account) Check(uidAccount string) (status bool) {
	if a.uid != uidAccount {
		fmt.Println("Аккаунт не найден")
		return false
	}

	fmt.Println("Аккаунт подтвержден")
	return true
}

// Accept accepts the visitor
func (l *account) Accept(v visitor.Visitor) (err error) {
	return v.VisitFoAccount(l)
}

// NewList creates an instance of the List
func NewAccount() Accounter {
	return &account{
		uid:     "user1",
		balance: 100,
	}
}
