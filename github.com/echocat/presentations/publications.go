package presentations

import (
	"fmt"

	"github.com/ariel17/katas/github.com/echocat/models"
)

type Printable interface {
	Print() string
}

func MergePrintables(books []models.Book, magazines []models.Magazine) []Printable {
	printables := []Printable{}
	for _, b := range books {
		printables = append(printables, b)
	}
	for _, m := range magazines {
		printables = append(printables, m)
	}
	return printables
}

func ShowAll(publications []Printable) {
	for _, p := range publications {
		fmt.Printf("+ %s\n", p.Print())
	}
}
