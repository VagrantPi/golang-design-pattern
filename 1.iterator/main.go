package main

import (
	"fmt"

	"design_pattern/1.iterator/book"
	"design_pattern/1.iterator/model"
)

func main() {
	var bookShelf model.Aggregate
	b := &model.BookShelf{}
	b.BookShelf(3)
	bookShelf = b
	aBook := &book.Book{}
	aBook.SetName("Around the World in 80 Days")
	bookShelf.Append(*aBook)
	bBook := &book.Book{}
	bBook.SetName("Bible")
	bookShelf.Append(*bBook)
	cBook := &book.Book{}
	cBook.SetName("Cinderella")
	bookShelf.Append(*cBook)
	dBook := &book.Book{}
	dBook.SetName("Daddy-Long-Legs")
	bookShelf.Append(*dBook)
	it := bookShelf.Iterator()
	for it.HasNext() {
		book := it.Next().(book.Book)
		fmt.Println(book.GetName())
	}
}
