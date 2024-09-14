package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    os.MkdirAll("./uploads", os.ModePerm)
    http.HandleFunc("/upload", uploadHandler)

    http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

    fmt.Println("Server started on :8080")
    http.ListenAndServe("localhost:8080", nil)
}