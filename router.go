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

var routingCollection = [50]routeMapStruct {
  // 1
  routeMapStruct{
    urlRoute: "/login",
    routeFunction: controllers.Login,
    method: "POST",
  },
  // 4
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
  // 4
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
  // 5
  routeMapStruct{
    urlRoute: "/order",
    routeFunction: controllers.ShowOrders,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/order/{id}",
    routeFunction: controllers.GetOrder,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/order/{id}",
    routeFunction: controllers.CreateOrder,
    method: "POST",
  },
  routeMapStruct{
    urlRoute: "/order/{id}",
    routeFunction: controllers.UpdateOrder,
    method: "PUT",
  },
  routeMapStruct{
    urlRoute: "/order/{id}",
    routeFunction: controllers.DeleteOrder,
    method: "DELETE",
  },
  // 5
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
  // 4
  routeMapStruct{
    urlRoute: "/outward_invoice",
    routeFunction: controllers.ShowOutwardInvoices,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/outward_invoice/{id}",
    routeFunction: controllers.GetOutwardInvoice,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/outward_invoice/{id}",
    routeFunction: controllers.CreateOutwardInvoice,
    method: "POST",
  },
  routeMapStruct{
    urlRoute: "/outward_invoice/{id}",
    routeFunction: controllers.UpdateOutwardInvoice,
    method: "PUT",
  },
  // 4
  routeMapStruct{
    urlRoute: "/outward_invoice_detail",
    routeFunction: controllers.ShowOutwardInvoiceDetails,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/outward_invoice_detail/",
    routeFunction: controllers.GetOutwardInvoiceDetail,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/outward_invoice_detail",
    routeFunction: controllers.CreateOutwardInvoiceDetail,
    method: "POST",
  },
  routeMapStruct{
    urlRoute: "/outward_invoice_detail",
    routeFunction: controllers.UpdateOutwardInvoiceDetail,
    method: "PUT",
  },
  // 5
  routeMapStruct{
    urlRoute: "/procure",
    routeFunction: controllers.ShowProcures,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/procure/{id}",
    routeFunction: controllers.GetProcure,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/procure/{id}",
    routeFunction: controllers.CreateProcure,
    method: "POST",
  },
  routeMapStruct{
    urlRoute: "/procure/{id}",
    routeFunction: controllers.UpdateProcure,
    method: "PUT",
  },
  routeMapStruct{
    urlRoute: "/procure/{id}",
    routeFunction: controllers.DeleteProcure,
    method: "DELETE",
  },
  // 5
  routeMapStruct{
    urlRoute: "/seller",
    routeFunction: controllers.ShowSellers,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/seller/{id}",
    routeFunction: controllers.GetSeller,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/seller/{id}",
    routeFunction: controllers.CreateSeller,
    method: "POST",
  },
  routeMapStruct{
    urlRoute: "/seller/{id}",
    routeFunction: controllers.UpdateSeller,
    method: "PUT",
  },
  routeMapStruct{
    urlRoute: "/seller/{id}",
    routeFunction: controllers.DeleteSeller,
    method: "DELETE",
  },
  // 4
  routeMapStruct{
    urlRoute: "/storefront",
    routeFunction: controllers.ShowStorefronts,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/storefront/",
    routeFunction: controllers.GetStorefront,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/storefront",
    routeFunction: controllers.CreateStorefront,
    method: "POST",
  },
  routeMapStruct{
    urlRoute: "/storefront",
    routeFunction: controllers.UpdateStorefront,
    method: "PUT",
  },
  // 5
  routeMapStruct{
    urlRoute: "/user",
    routeFunction: controllers.ShowUsers,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/user/{id}",
    routeFunction: controllers.GetUser,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/user/{id}",
    routeFunction: controllers.CreateUser,
    method: "POST",
  },
  routeMapStruct{
    urlRoute: "/user/{id}",
    routeFunction: controllers.UpdateUser,
    method: "PUT",
  },
  routeMapStruct{
    urlRoute: "/user/{id}",
    routeFunction: controllers.DeleteUser,
    method: "DELETE",
  },
  // 4
  routeMapStruct{
    urlRoute: "/warehouse",
    routeFunction: controllers.ShowWarehouses,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/warehouse/",
    routeFunction: controllers.GetWarehouse,
    method: "GET",
  },
  routeMapStruct{
    urlRoute: "/warehouse",
    routeFunction: controllers.CreateWarehouse,
    method: "POST",
  },
  routeMapStruct{
    urlRoute: "/warehouse",
    routeFunction: controllers.UpdateWarehouse,
    method: "PUT",
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
