package visitor

type accounter interface {
	Reserved(uidAccount string) (status bool)
	Accept(v Visiter)
}

// Интерфейс визитора реализующий функции для account
type Visiter interface {
	VisitFoAccount(a accounter)
}
