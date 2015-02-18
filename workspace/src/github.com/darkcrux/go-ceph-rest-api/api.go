package main

import (
	"log"

	"net/http"

	"github.com/AcalephStorage/grados"
	"github.com/gorilla/mux"
)

const (
	apiVersion = "0.1"
)

var (
	ceph *grados.Cluster

	uriRoot = "/api/v" + apiVersion

	// using a map for routes and functions...
	// this may change if it doesn't work for other impl.
	routes = map[string]func(res http.ResponseWriter, req *http.Request){
		uriRoot + "/fsid":              fsidHandler,
		uriRoot + "/osd/pool/{action}": poolHandler,
	}
)

func main() {
	// create connection to grados
	cluster, err := grados.ConnectToDefaultCluster()
	if err != nil {
		log.Fatalln("Ceph Cluster Error:", err)
	} else {
		ceph = cluster
		log.Println("Connected to Ceph.")
	}

	// adding the routes
	r := mux.NewRouter()
	for k, v := range routes {
		r.HandleFunc(k, v)
	}
	http.Handle("/", r)

	log.Println("ceph api started.")
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", nil))
}
