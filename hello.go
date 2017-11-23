package main

import (
    "fmt"
    "net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hello World\n")
}

func main() {
    http.HandleFunc("/", indexPage)
    fmt.Println("Start! Post :8080")
    http.ListenAndServe(":8080", nil)
    //fmt.Printf("hello, world\n")
}