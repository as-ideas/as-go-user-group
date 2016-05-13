package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"html/template"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func shortenerHandler(w http.ResponseWriter, r *http.Request) {
	my_map := make(map[string]string)
	
	
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("./login.gtpl")
        t.Execute(w, nil)
    }else {
        r.ParseForm()
       
       // fmt.Println("username:", r.Form["url"])
		my_url := r.Form["url"][0]
		my_map[my_url] = "test"
		io.WriteString(w, "The url you entered is "+my_url)
    }
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/shortener", shortenerHandler)

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)

	//http.ListenAndServe(":8000", nil)
}
