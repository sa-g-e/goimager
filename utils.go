package main

import (
    "mime/multipart"
    "path/filepath"
    "strings"
)

func isImage(fileHeader *multipart.FileHeader) bool {
    fileExt := strings.ToLower(filepath.Ext(fileHeader.Filename))
    allowedExtensions := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".bmp": true}
    return allowedExtensions[fileExt]
}