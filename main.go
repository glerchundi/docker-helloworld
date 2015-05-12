package main

import (
  _ "github.com/codegangsta/envy/autoload"
  "github.com/codegangsta/envy/lib"
  "github.com/codegangsta/negroni"
  "github.com/julienschmidt/httprouter"
  "net/http"
  "os"
)

func main() {
  n := negroni.Classic()
  r := httprouter.New()

  r.GET("/", HelloWorld)

  n.UseHandler(r)
  n.Run(":" + envy.MustGet("PORT"))
}

func HelloWorld(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
  hostname, _ := os.Hostname()
  rw.Write([]byte("Hello from " + hostname))
}
