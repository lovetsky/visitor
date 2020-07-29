package addons

import (
	"fmt"

	"github.com/labstack/gommon/color"
)

type accounter interface {
	Reserved(uidAccount string) (status bool)
	Accept(v Visiter)
}

// Интерфейс визитора реализующий функции для account
type Visiter interface {
	VisitFoAccount(int2 int)
}

type calcProduct struct {
	count int
}

// VisitFoAccount реализует вызов для объекта количества товара
func (c *calcProduct) VisitFoAccount(a int) {
	done := color.Green("[ OK ]")
	fmt.Println(done, " Добавлено товаров (новый метод): ", c.count)
}

// NewAccount создает реализацию интерефейса accounter
func NewVisitor(count int) Visiter {
	return &calcProduct{
		count: count,
	}
}
