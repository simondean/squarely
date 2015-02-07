package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
)

var templates = template.Must(template.ParseFiles("tmpl/dashboard.html"))

type dashboard struct {
	value int
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling request\n")
	d := new(dashboard)

	err := templates.ExecuteTemplate(w, "dashboard.html", d)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type randomCollector struct {
	lastValue int
}

func collect(collector *randomCollector) int {
	delta := rand.Intn(10) - 5
	nextValue := collector.lastValue + delta
	if nextValue < 0 || nextValue > 100 {
		nextValue = collector.lastValue - delta
	}

	collector.lastValue = nextValue
	return nextValue
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	//collector := new(randomCollector)

	//for {
	//	fmt.Printf("%v\n", collect(collector))
	//}
}
