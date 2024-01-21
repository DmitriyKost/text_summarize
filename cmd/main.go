package main

import (
    "github.com/DmitriyKost/text_summarize/pkg"
    "net/http"
)

func main() {
    http.HandleFunc("/", pkg.Index)
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/api/", http.StripPrefix("/api/", fs))
    http.HandleFunc("/summarize", pkg.Summarize)
    http.ListenAndServe(":8080", nil)
}
