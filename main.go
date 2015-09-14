package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/healthcheck", healthcheckHandler)

  log.Fatal(http.ListenAndServe(":8181", nil))
}


func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "WORKING")
}
