package model

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type BookShelfIterator struct {
	bookShelf BookShelf
	index     int
}

func (b *BookShelfIterator) BookShelfIterator(bookShelf *BookShelf) BookShelfIterator {
	b.bookShelf = *bookShelf
	b.index = 0
	return *b
}

func (b *BookShelfIterator) HasNext() bool {
	if b.index < b.bookShelf.GetLen() {
		return true
	}
	return false
}

func (b *BookShelfIterator) Next() interface{} {
	book := b.bookShelf.GetBookAt(b.index)
	b.index++
	return book
}
