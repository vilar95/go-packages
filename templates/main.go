package main

import (
	"html/template"
	"os"
	"strings"
)


type Course struct {
	Name        string
	Workload    int
}

type Courses []Course

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

		t := template.New("content.html")
		t.Funcs(template.FuncMap{
			"ToUpper": strings.ToUpper,})
		t = template.Must(t.ParseFiles(templates...))
		err := t.Execute(os.Stdout, Courses{
			{"Go Programming", 40},
			{"Python Basics", 30},
			{"Flutter Development", 50},
			{"JavaScript Essentials", 45},
			{"Assembly Language", 60},
			{"C++ Advanced", 55},
			{"Cobol Fundamentals", 35},
			{"Vbasic Programming", 25},
			{"C# Intermediate", 50},
			{"Linux System Administration", 70},
		})
		if err != nil {
			panic(err)
		}
}