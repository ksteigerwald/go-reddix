package store

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

var controller = &Controller{ Repository: Repository{} }

// Route defines a route
type Route struct {
  Name string
  Method string
  Pattern string
  HandlerFunc http.HandlerFunc
}

// Routes dfines the list of routes
type Routes []Route

var routes = Routes {
  Route {
    "Authentication",
    "POST",
    "/get-token",
    controller.GetToken,
  },
  Route {
    "CreateThread",
    "POST",
    "/threads",
    controller.CreateObj,
  },
}

func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)
  for _, route := range routes {
    var handler http.Handler
    log.Println(route.Name)
    handler = route.HandlerFunc

    router.Methods(route.Method).
    Path(route.Pattern).
    Name(route.Name).
    Handler(handler)
  }

  return router
}
