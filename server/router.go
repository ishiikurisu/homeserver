package server

import "net/http"
import "fmt"

type ActionRouter struct {
  Actions map[string]func()string
}

func NewRouter() *ActionRouter {
  router := ActionRouter {
    Actions: make(map[string]func()string)
  }
  return &router
}

func (router *ActionRouter) Route(method string, action func()string) {
  router.Actions[method] = action
  http.HandleFunc(method, func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s\n", action())
  })
}

func (router *ActionRouter) Serve(port string) {
  http.ListenAndServe(":" + port, nil)
}
