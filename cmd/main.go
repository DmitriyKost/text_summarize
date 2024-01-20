package main

import (
    "github.com/DmitriyKost/text_summarize/pkg"
    "net/http"
)

func main() {
    http.HandleFunc("/", pkg.Index)
    http.ListenAndServe(":8080", nil)
}
