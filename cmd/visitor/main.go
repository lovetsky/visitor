// Пакет является примером реализации паттерна Visitor
package main

import (
	"github.com/lovetsky/visitor/internal/account"
	"github.com/lovetsky/visitor/internal/addons"
)

func main() {
	f := addons.NewCalcProduct(10)
	a := account.NewAccount("user1", f)
	a.Reserved("user1")
	a.Accept(f)

	return
}
