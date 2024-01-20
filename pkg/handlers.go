package pkg

import (
	"fmt"
	"io"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "index.html", nil)
}

func Summarize(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
    }
    defer r.Body.Close()
    text := string(body)

    // handling text TODO!!!
    fmt.Printf("%s", text)

    w.WriteHeader(http.StatusOK)
    // return summarized text TODO!!!
    w.Write([]byte("Text received"))
}
