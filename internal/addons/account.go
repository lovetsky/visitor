package addons

import (
	"fmt"
	"github.com/labstack/gommon/color"
)

var done = color.Green("[ OK ]")

type calcProduct struct {
	count int
}

// Интерфейс визитора реализующий функции для account
type Visiter interface {
	VisitFoAccount(int2 int)
}

type accounter interface {
	Reserved(uidAccount string) (status bool)
	Accept(v Visiter)
}

// VisitFoAccount реализует вызов для объекта количества товара
func (c *calcProduct) VisitFoAccount(a int) {
	fmt.Println(done, " Добавлено товаров (новый метод): ", c.count)
}

// NewAccount создает реализацию интерефейса accounter
func NewCalcProduct(count int) Visiter {
	return &calcProduct{
		count: count,
	}
}
