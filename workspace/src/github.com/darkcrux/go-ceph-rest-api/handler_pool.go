package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
)

const (
  // actions
  create = "create"
  // get = "get"
  // delete = "delete"

  // methods
  // get = "GET"
  // post = "POST"
  put = "PUT"
  // delete = "DELETE"
)

type PoolAPI struct {
  Res http.ResponseWriter
  Req *http.Request
  Params map[string]string
}

type Response struct {
  Status string      `json:"status"`
  Output interface{} `json:"output"`
}

func poolHandler(res http.ResponseWriter, req *http.Request) {
  res.WriteHeader(200)
  var pool PoolAPI
  pool.Res = res
  pool.Req = req
  pool.Params = mux.Vars(req)
  callback := func(resp Response) {
    json_response, _ := json.Marshal(resp)
    fmt.Fprint(res, string(json_response))
  }
  pool.handleAction(callback)
}

func (pool PoolAPI) handleAction(callback func(resp Response)) {
  action := pool.Params["action"]
  switch {
  case action == create && pool.Req.Method == put:
    resp := pool.create()
    callback(resp)
  }
}

func (pool PoolAPI) create() (resp Response) {
  poolName := pool.Req.URL.Query().Get("pool")
  err := ceph.CreatePool(poolName)
  if err != nil {
    resp.Status = err.Error()
  } else {
    resp.Status = "pool '" + poolName + "' created."
  }
  resp.Output = []string{}
  return
}

// add other CRUD methods here...
