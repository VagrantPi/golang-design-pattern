package model

import (
	"design_pattern/1.iterator/model/book"
)

// Aggregate - operation interface collection
type Aggregate interface {
	Iterator() Iterator
	// 這邊會多個 Append 為了方便 Aggregate 的 instance(bookShelf) 可以直接使用 Append 這個 Method
	// Iterator Pattern 並不存在 Append 的 interface
	Append(interface{})
}

// ========= other BookShelf =========

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

// func (b *BookShelf) Append(abook interface{}) {
// 	if len(b.books) == 0 {
// 		b.books = abook.(book.Book).GetName()
// 	} else {
// 		b.books = b.books + "," + abook.(book.Book).GetName()
// 	}
// }

// func (b *BookShelf) GetLen() int {
// 	s := strings.Split(b.books, ",")
// 	return len(s)
// }

// func (b *BookShelf) Iterator() Iterator {
// 	i := &BookShelfIterator{}
// 	i.BookShelfIterator(b)
// 	return i
// }

// ===================================

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

func (b *BookShelf) Append(abook interface{}) {
	if b.GetLen() >= len(b.books) {
		newBooks := make([]book.Book, b.last*2)
		for index, item := range b.books {
			newBooks[index] = item
		}
		b.books = newBooks
	}
	b.books[b.last] = abook.(book.Book)
	b.last++
}

func (b *BookShelf) GetLen() int {
	return b.last
}

func (b *BookShelf) Iterator() Iterator {
	i := &BookShelfIterator{}
	i.BookShelfIterator(b)
	return i
}
