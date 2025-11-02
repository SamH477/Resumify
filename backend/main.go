package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from Go!")
    })

    fmt.Println("Go server running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
