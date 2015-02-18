package main

import (
	"fmt"

	"net/http"

	"encoding/json"
	"github.com/gorilla/mux"
)

const (
	// actions
	ActionCreate = "create"
	// ActionGet = "get"
	// ActionDelete = "delete"

	// methods
	// MethodGet = "GET"
	// MethodPost = "POST"
	MethodPut = "PUT"
	// MethodDelete = "DELETE"
)

type poolAPI struct {
	Res    http.ResponseWriter
	Req    *http.Request
	Params map[string]string
}

type Response struct {
	Status string      `json:"status"`
	Output interface{} `json:"output"`
	Code   int         `json:code`
}

func poolHandler(res http.ResponseWriter, req *http.Request) {
	pool := poolAPI{res, req, mux.Vars(req)}
	callback := func(resp Response) {
		res.WriteHeader(resp.Code)
		json_response, _ := json.Marshal(resp)
		fmt.Fprint(res, string(json_response))
	}

	switch req.Method {
	default:
		res.WriteHeader(204)               // Default; No Content
		fmt.Fprint(res, "204: No Content") // Will replace this later
	case MethodPut:
		pool.handlePut(callback)
		// case MethodPost:
		// 	pool.handlePost(callback)
		// case MethodGet:
		// 	pool.handleGet(callback)
		// case MethodDelete:
		// 	pool.handleDelete(callback)
	}
}

func (pool poolAPI) handlePut(callback func(resp Response)) {
	switch pool.Params["action"] {
	case ActionCreate:
		resp := pool.create()
		callback(resp)
	}
}

func (pool poolAPI) create() (resp Response) {
	poolName := pool.Req.URL.Query().Get("pool")
	if err := ceph.CreatePool(poolName); err != nil {
		resp.Status = err.Error()
		resp.Code = 400
	} else {
		resp.Status = fmt.Sprintf("pool '%s' created.", poolName)
		resp.Code = 201
	}
	resp.Output = []string{}
	return
}

// add other CRUD methods here...
