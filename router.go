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
  query map[string]string
}

var routingCollection = [6]routeMapStruct {
  routeMapStruct{
    urlRoute: "/login",
    routeFunction: controllers.Login,
    method: "POST",
  },
  routeMapStruct{
    urlRoute: "/product",
    routeFunction: controllers.ShowProducts,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/product/{code}",
    routeFunction: controllers.GetProduct,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/product/{code}",
    routeFunction: controllers.CreateProduct,
    method: "POST",
  },
  routeMapStruct{
    urlRoute: "/product/{code}",
    routeFunction: controllers.UpdateProduct,
    method: "PUT",
  },
  routeMapStruct{
    urlRoute: "/product/{code}",
    routeFunction: controllers.DeleteProduct,
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
