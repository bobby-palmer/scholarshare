package main

import (
	"log"
	"net/http"
	"strconv"
)

const PORT = 8080

func main() {
  mux := http.NewServeMux()

  mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("frontend/build"))))

  err := http.ListenAndServe(":" + strconv.FormatUint(PORT, 10), mux)
  if err != nil {
    log.Fatal(err)
  }
}
