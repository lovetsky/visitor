package main

import (
	"fmt"
	"github.com/labstack/gommon/color"
)

var done = color.Green("[ OK ]")
var fail = color.Red("[Fail]")

type accounter interface {
	Reserved(uidAccount string) (status bool)
	Accept(v visitor)
}

type visitor interface {
	VisitFoAccount(*account)
}

type account struct {
	uid     string
	balance int
}

type calcProduct struct {
	area int
}

// Reserved реализует метод резервирования товара покупателем
func (a *account) Reserved(uidAccount string) (status bool) {
	fmt.Println(done, " Товар был зарезервирован для покупателя: ", uidAccount)
	return true
}

// Accept
func (c *account) Accept(v visitor) {
	v.VisitFoAccount(c)
}

// NewAccount создает реализацию интерефейса accounter
func NewAccount(uidAccount string) accounter {
	return &account{
		uid:     uidAccount,
		balance: 100,
	}
}

// VisitFoAccount добавленме нового метода
func (a *calcProduct) VisitFoAccount(s *account) {
	fmt.Println(done, " Баланс пользователя (новый метод): ", s.balance)
}

func main() {
	v := NewAccount("user1")
	v.Reserved("user1")

	calcProduct := &calcProduct{}
	v.Accept(calcProduct)

	return
}
