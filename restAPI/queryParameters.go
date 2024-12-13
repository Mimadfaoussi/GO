package main

import (
	"fmt"
	"net/http"
)


func queryHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	name := query.Get("name")

	if name == "" {
		fmt.Fprintln(w, "Please provide a name")
	} else {
		fmt.Fprintf(w, "Hello, %s!\n", name)
	}
}

func main() {
	http.HandleFunc("/greet", queryHandler)
	fmt.Println("server is running ...")
	http.ListenAndServe(":8000", nil)
}
