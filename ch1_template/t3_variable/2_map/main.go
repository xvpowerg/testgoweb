package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.html"))
}
func main() {
	studentScore :=
		map[string]int32{
			"Ken":   85,
			"Vivin": 72,
			"Lindy": 87,
		}
	err := tpl.Execute(os.Stdout, studentScore)
	if err != nil {
		log.Fatalln(err)
	}
}
