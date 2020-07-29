// пакет реализует пример паттерна Visitor
package main

import (
	"github.com/lovetsky/visitor/pkg/account"
	"github.com/lovetsky/visitor/pkg/visitor"
)

func main() {
	l := account.NewAccount()
	l.Check("user1")

	arrL := visitor.NewVisitor(100)
	l.Accept(arrL)
}
