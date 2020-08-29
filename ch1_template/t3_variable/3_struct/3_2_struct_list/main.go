package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Student struct {
	Name string
	Age  int16
}

type Book struct {
	BookName string
	Price    int32
}

func init() {
	tpl = template.Must(template.ParseFiles("template.html"))
}

func main() {
	st1 := Student{
		"Ken",
		10,
	}

	st2 := Student{
		"Vivin",
		25,
	}

	book1 := Book{
		"Android",
		520,
	}
	book2 := Book{
		"IOS",
		650,
	}

	stList := []Student{st1, st2}
	bookList := []Book{book1, book2}
	studentStruct := struct {
		StudentList []Student
		BookList    []Book
	}{
		stList,
		bookList,
	}
	err := tpl.Execute(os.Stdout, studentStruct)
	if err != nil {
		log.Fatalln(err)
	}
}
