package main

import (
  "github.com/codegangsta/negroni"
  "github.com/julienschmidt/httprouter"
  "net/http"
  "os"
  "time"
)

func main() {
  ds := os.Getenv("DELAY")
  if ds != "" {
    if d, err := time.ParseDuration(ds); err == nil {
      time.Sleep(d)
    }
  }
  p := os.Getenv("PORT")
  if p == "" {
    p = "8080"
  }
  n := negroni.Classic()
  r := httprouter.New()

  r.GET("/", HelloWorld)

  n.UseHandler(r)
  n.Run(":" + p)
}

func HelloWorld(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
  hostname, _ := os.Hostname()
  rw.Write([]byte("Hello from " + hostname))
}
