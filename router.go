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

var routingCollection = [15]routeMapStruct {
  routeMapStruct{
    urlRoute: "/login",
    routeFunction: controllers.Login,
    method: "POST",
  },
  routeMapStruct{
    urlRoute: "/inward_invoice",
    routeFunction: controllers.ShowInwardInvoices,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/inward_invoice/{invoice_code}",
    routeFunction: controllers.GetInwardInvoice,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/inward_invoice/{invoice_code}",
    routeFunction: controllers.CreateInwardInvoice,
    method: "POST",
  },
  routeMapStruct{
    urlRoute: "/inward_invoice/{invoice_code}",
    routeFunction: controllers.UpdateInwardInvoice,
    method: "PUT",
  },
  routeMapStruct{
    urlRoute: "/inward_invoice/{invoice_code}",
    routeFunction: controllers.DeleteInwardInvoice,
    method: "DELETE",
  },
  routeMapStruct{
    urlRoute: "/inward_invoice_detail",
    routeFunction: controllers.ShowInwardInvoiceDetails,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/inward_invoice_detail/",
    routeFunction: controllers.GetInwardInvoiceDetail,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/inward_invoice_detail",
    routeFunction: controllers.CreateInwardInvoiceDetail,
    method: "POST",
  },
  routeMapStruct{
    urlRoute: "/inward_invoice_detail",
    routeFunction: controllers.UpdateInwardInvoiceDetail,
    method: "PUT",
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
