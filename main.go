package main

import (
  "github.com/joho/godotenv"
  "fmt"
)

// our main function
func main() {
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading .env file")
    return
  }

  migrate()
  startRouter(":6789")
}
