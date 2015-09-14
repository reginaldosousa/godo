package main

import (
  "log"
  "net/http"

  "github.com/reginaldosousa/godo/handlers"
)

func main() {
  http.HandleFunc("/healthcheck", handlers.HealthcheckHandler)

  log.Fatal(http.ListenAndServe(":8181", nil))
}
