package routes

import(
  "fmt"
  "net/http"
  "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter,r *http.Request, params httprouter.Params) {
  fmt.Fprint(w,"Welcome to golang!\n")
}
