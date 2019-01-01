package model

import (
	"design_pattern/1.iterator/book"
)

type Aggregate interface {
	Interator() Interator
}

// type BookShelf struct {
// 	books string
// 	last  int
// }

// func (b *BookShelf) BookShelf(maxsize int) {
// 	// b.books = make([]book.Book, maxsize)
// }

// func (b *BookShelf) GetBookAt(index int) book.Book {
// 	s := strings.Split(b.books, ",")
// 	aBook := book.Book{}
// 	aBook.SetName(s[index])
// 	return aBook
// }

// func (b *BookShelf) AppendBook(abook book.Book) {
// 	if len(b.books) == 0 {
// 		b.books = abook.GetName()
// 	} else {
// 		b.books = b.books + "," + abook.GetName()
// 	}
// }

// func (b *BookShelf) GetLen() int {
// 	s := strings.Split(b.books, ",")
// 	return len(s)
// }

// func (b *BookShelf) Interator() BookShelfIterator {
// 	i := &BookShelfIterator{}
// 	return i.BookShelfIterator(b)
// }

type BookShelf struct {
	books []book.Book
	last  int
}

func (b *BookShelf) BookShelf(maxsize int) {
	b.books = make([]book.Book, maxsize)
}

func (b *BookShelf) GetBookAt(index int) book.Book {
	return b.books[index]
}

func (b *BookShelf) AppendBook(abook book.Book) {
	if b.GetLen() >= len(b.books) {
		newBooks := make([]book.Book, b.last*2)
		for index, item := range b.books {
			newBooks[index] = item
		}
		b.books = newBooks
	}
	b.books[b.last] = abook
	b.last++

}

func (b *BookShelf) GetLen() int {
	return b.last
}

func (b *BookShelf) Interator() BookShelfIterator {
	i := &BookShelfIterator{}
	return i.BookShelfIterator(b)
}
