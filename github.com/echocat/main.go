package main

import (
	"github.com/ariel17/katas/github.com/echocat/models"
	"github.com/ariel17/katas/github.com/echocat/presentations"
)

func main() {

	_, err := models.LoadAuthors("./authors.csv")
	if err != nil {
		panic(err)
	}

	books, err := models.LoadBooks("./books.csv")
	if err != nil {
		panic(err)
	}

	magazines, err := models.LoadMagazines("./magazines.csv")
	if err != nil {
		panic(err)
	}

	printables := presentations.MergePrintables(books, magazines)
	presentations.ShowAll(printables)
}