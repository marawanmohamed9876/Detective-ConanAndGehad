package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template
var locations map[string]int

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	locations = make(map[string]int)
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fname := r.FormValue("firster")

	loc := fname
	_, ok := locations[loc]
	fmt.Println(ok)
	/* if ok is true, entry is present otherwise entry is absent*/
	if ok {
		locations[loc] = locations[loc] + 1
	} else {
		locations[loc] = 1
	}

	for l := range locations {
		fmt.Println("The frequency of", l, "is", locations[l])
	}

	fmt.Println("--------------------------------------")
	max := 0
	maxPlace := ""
	for l := range locations {
		if max < locations[l] {
			max = locations[l]
			maxPlace = l
		}
	}
	d := struct {
		First string
		Max   string
	}{
		First: fname,
		Max:   maxPlace,
	}

	fmt.Println(" the most place that most neighbors agrees on until now is ", maxPlace, "the occurenss of the place is ", max)
	tpl.ExecuteTemplate(w, "processor.gohtml", d)
	// if r.FormValue("red") {
	// 	http.Redirect(w, r, "/index.gohtml", http.StatusSeeOther)
	// }

}
