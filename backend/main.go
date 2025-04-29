package main

import (
	"log"
	"net/http"
	"os"
)

const uploadDir = "./uploads"
const port = "8080"

// Create uploads dir at boot if not already made
func init() {
  if err := os.MkdirAll(uploadDir, 0755); err != nil {
    log.Fatal(err)
  }
}

func main() {
  http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir(uploadDir))))
  http.HandleFunc("/ls", getPDFListHandler)

  http.ListenAndServe(":" + port, nil)
}
