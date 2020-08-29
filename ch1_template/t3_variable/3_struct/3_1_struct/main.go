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

func init() {
	tpl = template.Must(template.ParseFiles("template.html"))
}

func main() {
	studentStruct := Student{
		Name: "Ken",
		Age:  10,
	}
	err := tpl.Execute(os.Stdout, studentStruct)
	if err != nil {
		log.Fatalln(err)
	}
}
