package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}
type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var people []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for _, item := range people {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Person{})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  var person Person
  _ = json.NewDecoder(r.Body).Decode(&person)
  person.ID = params["id"]
  people = append(people, person)
  json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  var newPeople []Person
  for _, person := range people {
    if person.ID != params["id"] {
      newPeople = append(newPeople, person)
    }
  }
  people = newPeople
  json.NewEncoder(w).Encode(people)
}

// our main function
func main() {
  startRouter(":6789")
}
