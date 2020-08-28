package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("my.html")
	if err != nil {
		log.Fatalln(err)
	}
	//輸出到Terminal
	//err = tpl.Execute(os.Stdout, nil)
	//輸出到指定目錄
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
	//一次ParseFiles 多筆
	// tpl, err = template.ParseFiles("two.html", "Three.html")
	//可限定Parse的類型
	//tpl, err = template.ParseGlob("*")
	//如果ParseGlob 有錯誤Must會拋出panic
	tpl = template.Must(template.ParseGlob("*.html"))
	tpl.ExecuteTemplate(os.Stdout, "two.html", nil)
	tpl.ExecuteTemplate(os.Stdout, "Three.html", nil)
	//取得第一筆
	tpl.Execute(os.Stdout, nil)
}
