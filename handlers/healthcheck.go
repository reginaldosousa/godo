package handlers

import (
  "fmt"
  "net/http"
)

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "WORKING")
}
