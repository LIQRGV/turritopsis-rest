package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "turritopsis-rest/controllers"
)

type routeMapStruct struct {
  urlRoute string
  routeFunction func(http.ResponseWriter, *http.Request)
  method string
}

var routingCollection = [4]routeMapStruct{
  routeMapStruct{
    urlRoute: "/login",
    routeFunction: controllers.Login,
    method: "POST",
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
