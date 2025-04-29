package main

import (
	"net/http"
	"strconv"
)

const PORT = 8080

func main() {
  mux := http.NewServeMux()
  http.ListenAndServe(":" + strconv.FormatUint(PORT, 10), mux)
}
