package main

import (
	"log"
	"net/http"
	"strconv"
)

const PORT = 8080

func main() {
  mux := http.NewServeMux()

  err := http.ListenAndServe(":" + strconv.FormatUint(PORT, 10), mux)
  if err != nil {
    log.Fatal(err)
  }
}
