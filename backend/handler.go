package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func uploadPDFHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    http.Error(w, "Bad Method", http.StatusInternalServerError)
    return
  }

  r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

  err := r.ParseMultipartForm(10 << 20)
  if err != nil {
      http.Error(w, "Could not parse multipart form", http.StatusBadRequest)
      return
  }

  file, fileHeader, err := r.FormFile("pdf")
  if err != nil {
      http.Error(w, "File not found in form", http.StatusBadRequest)
      return
  }
  defer file.Close()

  if !strings.HasSuffix(fileHeader.Filename, ".pdf") {
      http.Error(w, "Only .pdf files allowed", http.StatusBadRequest)
      return
  }

  dstPath := filepath.Join(uploadDir, fileHeader.Filename)
  out, err := os.Create(dstPath)
  if err != nil {
      http.Error(w, "Could not save file", http.StatusInternalServerError)
      return
  }
  defer out.Close()

  _, err = io.Copy(out, file)
  if err != nil {
      http.Error(w, "Failed to write file", http.StatusInternalServerError)
      return
  }

  fmt.Fprintf(w, "Uploaded: %s\n", fileHeader.Filename)
}

func getPDFListHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet {
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
  }

  files, err := os.ReadDir(uploadDir)

  if err != nil {
    http.Error(w, "Fail to read pdfs", http.StatusInternalServerError)
    return
  }

  pdfs := make([]string, 0)

  for _, file := range files {
    if !file.IsDir() {
      pdfs = append(pdfs, file.Name())
    }
  }

  w.Header().Set("Content-Type", "application/json")
  err = json.NewEncoder(w).Encode(pdfs)

  if err != nil {
    http.Error(w, "Fail to encode file names", http.StatusInternalServerError)
  }
}
