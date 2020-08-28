package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.html"))
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "template.html", 42)
	if err != nil {
		fmt.Println(err)
	}
}
