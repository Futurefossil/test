package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/lock", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "Locked") })
    http.HandleFunc("/unlock", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "Unlocked") })
    http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "Status: locked/unlocked") })

    fmt.Println("Server started on port 8080")
    http.ListenAndServe(":8080", nil)
}
