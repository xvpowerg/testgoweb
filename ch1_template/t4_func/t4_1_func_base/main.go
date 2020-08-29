package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"up": toUp,
	"ft": three,
	"tf": monthDayYear}

func toUp(s string) string {
	s = strings.ToUpper(s)
	return s
}
func three(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}
func monthDayYear(t time.Time) string {
	return t.Format("2006-01-02")
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("template.html"))
}

func main() {
	// tpl = tpl.Funcs(fm)
	st := struct {
		Name string
		Time time.Time
	}{
		"howard",
		time.Now(),
	}
	err := tpl.ExecuteTemplate(os.Stdout, "template.html", st)
	if err != nil {
		log.Fatalln(err)
	}
}
