package main

import (
	"fmt"

	"design_pattern/1.iterator/book"
	"design_pattern/1.iterator/model"
)

func main() {
	bookShelf := &model.BookShelf{}
	bookShelf.BookShelf(3)
	aBook := &book.Book{}
	aBook.SetName("Around the World in 80 Days")
	bookShelf.AppendBook(*aBook)
	bBook := &book.Book{}
	bBook.SetName("Bible")
	bookShelf.AppendBook(*bBook)
	cBook := &book.Book{}
	cBook.SetName("Cinderella")
	bookShelf.AppendBook(*cBook)
	dBook := &book.Book{}
	dBook.SetName("Daddy-Long-Legs")
	bookShelf.AppendBook(*dBook)
	it := bookShelf.Interator()
	for it.HasNext() {
		book := it.Next().(book.Book)
		fmt.Println(book.GetName())
	}
}
