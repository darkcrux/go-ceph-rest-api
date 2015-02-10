# go-ceph-rest-api

Our pathetic attempt to port the Ceph API to go.

## Setup Workspace

We need access to ceph and the ceph-dev libraries to work with this. Boot up the vagrant box to start the workspace.

```
$ vagrant up
```

The workspace should be mounted on `/vagrant/workspace`. This is set as the `$GOPATH` to work on inside the VM. 

## Build Project

Building the binary is done inside the VM.

```
$ vagrant ssh
$ cd /vagrant/workspace/src/github.com/darkcrux/go-ceph-rest-api
$ go get
$ go build
```

## Running the API

At the moment, the binary will use the default ceph configuration. This needs to be updated in the future, but for now, run the app as su:

```
$ sudo ./go-ceph-rest-api
```

Once running, the API should be accessible via port 9000 (Needs to be updated).

```
$ curl http://200.200.200.200:9000/fsid
```

## Notes

Plenty of things to do:
- Create Makefile for the project
- Cross-compilation support
- Configuration options
- ???
