package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "strings"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    name := html.EscapeString(r.URL.Path)
    name = strings.Replace(name, "/", "", -1)
    fmt.Fprintf(w, "Fuck off, %s!!!", name)
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
