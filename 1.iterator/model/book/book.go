package book

type BookI interface {
	SetName(n string)
	GetName() string
}

type Book struct {
	name string
}

func (b *Book) SetName(n string) {
	b.name = n
}

func (b Book) GetName() string {
	return b.name
}
