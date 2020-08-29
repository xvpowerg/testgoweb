package main

import (
	"log"
	"os"
	"text/template"
)

var tp1 *template.Template

func init() {
	tp1 = template.Must(template.ParseFiles("template.html"))
}
func main() {
	names := []string{"Ken", "Vivin", "Join"}
	err := tp1.Execute(os.Stdout, names)
	if err != nil {
		log.Fatalln(err)
	}
}
