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
	st := struct {
		Names []string
		Value string
	}{
		names,
		"Title:",
	}
	err := tp1.Execute(os.Stdout, st)
	if err != nil {
		log.Fatalln(err)
	}
}
