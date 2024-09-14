package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strconv"
    "time"
)

const secretKey = "test" // Set your secret key here

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
        return
    }

    // Check for secret key in headers
    if r.Header.Get("X-Secret-Key") != secretKey {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    err := r.ParseMultipartForm(10 << 20) // 10 MB max upload size
    if err != nil {
        http.Error(w, "Could not parse form", http.StatusBadRequest)
        return
    }

    file, fileHeader, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Error retrieving file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    if !isImage(fileHeader) {
        http.Error(w, "Only image files are allowed", http.StatusUnsupportedMediaType)
        return
    }

    encodedTime := strconv.FormatInt(time.Now().UnixNano(), 10) // Get the Unix timestamp in nanoseconds

    savePath := "./uploads/" + encodedTime
    outFile, err := os.Create(savePath)
    if err != nil {
        http.Error(w, "Unable to save the file", http.StatusInternalServerError)
        return
    }
    defer outFile.Close()

    _, err = io.Copy(outFile, file)
    if err != nil {
        http.Error(w, "Error saving the file", http.StatusInternalServerError)
        return
    }

    fileURL := fmt.Sprintf("http://localhost:8080/uploads/%s", encodedTime)
    fmt.Fprintf(w, "{\"success\": true, \"url\": \"%s\"}\n", fileURL)
}
