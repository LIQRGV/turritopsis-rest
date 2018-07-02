package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type routeMapStruct struct {
  urlRoute string
  routeFunction func(http.ResponseWriter, *http.Request)
  method string
}

var routingCollection = [4]routeMapStruct{
  routeMapStruct{
    urlRoute: "/people",
    routeFunction: GetPeople,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/people/{id}",
    routeFunction: GetPerson,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/people/{id}",
    routeFunction: CreatePerson,
    method: "POST",
  },
  routeMapStruct{
    urlRoute: "/people/{id}", 
    routeFunction: DeletePerson,
    method: "DELETE",
  },
}

func createRouting(router *mux.Router) {
  for _, element := range routingCollection {
    router.HandleFunc(
      element.urlRoute, element.routeFunction,
    ).Methods(element.method)
  }
}

func startRouter(url string) {
  router := mux.NewRouter()
  createRouting(router)

  log.Fatal(http.ListenAndServe(url, router))
}
