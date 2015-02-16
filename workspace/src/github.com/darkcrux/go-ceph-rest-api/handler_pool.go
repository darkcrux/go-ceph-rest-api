package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

type PoolAPI struct {
  Res http.ResponseWriter
  Req *http.Request
}

func poolHandler(res http.ResponseWriter, req *http.Request) {
  res.WriteHeader(200)
  var pool PoolAPI
  pool.Res = res
  pool.Req = req
  callback := func(data string, err error) {
    fmt.Fprint(res, data)
  }
  pool.handleAction(callback)
}

func (pool PoolAPI) handleAction(callback func(data string, err error)) {
  vars := mux.Vars(pool.Req)
  action := vars["action"]
  switch action {
  case "create":
    data, err := pool.create()
    callback(data, err)
  }
}

// TODO: Consider methods POST, PUT, etc
func (pool PoolAPI) create() (data string, err error) {
  data = "creating..."
  err = nil
  return
}

// add other CRUD methods here...
