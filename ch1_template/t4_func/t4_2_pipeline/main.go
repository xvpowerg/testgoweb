package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

func pluse(x int) int {
	return x + x
}
func square(x int) float64 {
	return math.Pow(float64(x), 2)
}
func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"fp":    pluse,
	"fsq":   square,
	"fsqrt": sqRoot,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("template.html"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "template.html", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
