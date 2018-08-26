package controllers

import (
  "net/http"
)


func Login(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
}
