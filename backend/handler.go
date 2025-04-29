package main

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"
)

var mtx sync.Mutex

func uploadPDFHandler(w http.ResponseWriter, r *http.Request) {

}

func downloadPDFHandler(w http.ResponseWriter, r *http.Request) {

}

// List all PDFs availible
func getPDFListHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet {
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
  }

  mtx.Lock()
  files, err := os.ReadDir(uploadDir)
  mtx.Unlock()

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
