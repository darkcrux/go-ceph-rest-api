package main

import (
	"fmt"
	"net/http"
)

func fsidHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	fsid := ceph.FSID()
	fmt.Fprint(res, fsid)
}
