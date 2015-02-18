package main

import (
	"fmt"

	"net/http"

	"encoding/json"
	"github.com/gorilla/mux"
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

type poolAPI struct {
	Res    http.ResponseWriter
	Req    *http.Request
	Params map[string]string
}

type Response struct {
	Status string      `json:"status"`
	Output interface{} `json:"output"`
}

func poolHandler(res http.ResponseWriter, req *http.Request) {
	pool := poolAPI{res, req, mux.Vars(req)}
	callback := func(resp Response) {
		json_response, _ := json.Marshal(resp)
		fmt.Fprint(res, string(json_response))
	}
	pool.handlePut(callback)
	// pool.handlePost(callback)
	// pool.handleGet(callback)
	// pool.handleDelete(callback)
}

func (pool poolAPI) handlePut(callback func(resp Response)) {
	if action := pool.Params["action"]; pool.Req.Method == put {
		switch action {
		case create:
			resp := pool.create()
			callback(resp)
		}
	}
}

func (pool poolAPI) create() (resp Response) {
	poolName := pool.Req.URL.Query().Get("pool")
	if err := ceph.CreatePool(poolName); err != nil {
		resp.Status = err.Error()
		pool.Res.WriteHeader(400)
	} else {
		resp.Status = fmt.Sprintf("pool '%s' created.", poolName)
		pool.Res.WriteHeader(201)
	}
	resp.Output = []string{}
	return
}

// add other CRUD methods here...
