package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/hulk0301/example-app-backend/internal/router"
)

const port = 8080

func main() {
  http.HandleFunc("/", router.HandleRequest)
	log.Printf("Server listening on http://localhost:%d/", port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port) , nil))
}
