package main

import (
	"log"

	"net/http"

	"github.com/AcalephStorage/grados"
	"github.com/gorilla/mux"
)

var ceph *grados.Cluster

func main() {
	// create connection to grados
	cluster, err := grados.ConnectToDefaultCluster()
	if err != nil {
		log.Fatalln("Ceph Cluster Error:", err)
	} else {
		ceph = cluster
	}

	log.Println("ceph connection started")

	r := mux.NewRouter()
	r.HandleFunc("/fsid", fsidHandler)

	http.Handle("/", r)

	log.Println("ceph api started.")

	http.ListenAndServe("0.0.0.0:9000", nil)
}
