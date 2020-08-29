package main

import (
	"log"
	"os"
	"text/template"
)

var tp1 *template.Template

//Item struct
type Item struct {
	Name     string
	Des      string
	CanSales bool
}

func init() {
	tp1 = template.Must(template.ParseFiles("template.html"))
}
func main() {
	i1 := Item{
		Name:     "Android",
		Des:      "is a android",
		CanSales: false,
	}
	i2 := Item{
		Name:     "IPhone11",
		Des:      "is a IOS",
		CanSales: true,
	}
	i3 := Item{
		Name:     "",
		Des:      "is a IOS and Android",
		CanSales: true,
	}
	items := []Item{i1, i2, i3}
	err := tp1.Execute(os.Stdout, items)
	if err != nil {
		log.Fatalln(err)
	}
}
