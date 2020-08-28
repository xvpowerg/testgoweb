package main

import (
	"html/template"
	"log"
	"os"
)

var templ *template.Template

func init() {
	templ = template.Must(template.ParseFiles("template.html"))
}
func main() {
	names := []string{"Ken", "Vivin", "Lindy", "Jesus"}
	err := templ.Execute(os.Stdout, names)
	if err != nil {
		log.Fatalln(err)
	}

}
