# Type Template

## [template.Template](https://godoc.org/text/template#Template)
``` Go
template.Template
```

***

# Parsing templates

## [template.ParseFiles](https://godoc.org/text/template#ParseFiles)
``` Go
func ParseFiles(filenames ...string) (*Template, error)
```

## [template.ParseGlob](https://godoc.org/text/template#ParseGlob)
``` Go
func ParseGlob(pattern string) (*Template, error)
```
***

## [template.Parse](https://godoc.org/text/template#Template.Parse)
``` Go
func (t *Template) Parse(text string) (*Template, error)
```

## [template.ParseFiles](https://godoc.org/text/template#Template.ParseFiles)
``` Go
func (t *Template) ParseFiles(filenames ...string) (*Template, error)
//一次ParseFiles 多筆
tpl, err = template.ParseFiles("two.html", "Three.html")
```

## [template.ParseGlob](https://godoc.org/text/template#Template.ParseGlob)
``` Go
func (t *Template) ParseGlob(pattern string) (*Template, error)
//可限定Parse的類型
tpl, err = template.ParseGlob("*")
```

***

# Executing templates

## [template.Execute](https://godoc.org/text/template#Template.Execute)
``` Go
func (t *Template) Execute(wr io.Writer, data interface{}) error
//可指定Template的輸出
tpl.Execute(os.Stdout, nil)
```

## [template.ExecuteTemplate](https://godoc.org/text/template#Template.ExecuteTemplate)
``` Go
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
一次ParseFiles多筆時可指定要輸出的Template
tpl.ExecuteTemplate(os.Stdout, "two.html", nil)

```

***

# Helpful template functions 

## [template.Must](https://godoc.org/text/template#Must)
``` Go
func Must(t *Template, err error) *Template
```
``` Go
//如果ParseGlob 有錯誤Must會拋出panic
tpl = template.Must(template.ParseGlob("*"))
```

## [template.New](https://godoc.org/text/template#New)
``` Go
func New(name string) *Template
```

***

# The init function

## [The init function](https://golang.org/doc/effective_go.html#init)
``` Go
func init()
```
