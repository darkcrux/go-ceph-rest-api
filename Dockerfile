FROM ubuntu:14.04
MAINTAINER dexter.genterone@gmail.com

RUN apt-get update -y && \
    apt-get install -y wget && \
    wget -q -O- 'https://ceph.com/git/?p=ceph.git;a=blob_plain;f=keys/release.asc' | sudo apt-key add - && \
    echo deb http://ceph.com/debian-firefly/ $(lsb_release -sc) main | sudo tee /etc/apt/sources.list.d/ceph.list && \
    apt-get update -y && \
    apt-get install -y -q --no-install-recommends ceph librados-dev libcephfs-dev librbd-dev curl build-essential ca-certificates git mercurial bzr && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN mkdir /goroot && \
    mkdir /gopath && \
    curl https://storage.googleapis.com/golang/go1.4.1.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1

ENV GOROOT /goroot
ENV GOPATH /gopath
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

WORKDIR /gopath/src/go-ceph-rest-api
ADD workspace/src/github.com/darkcrux/go-ceph-rest-api /gopath/src/go-ceph-rest-api
RUN go get go-ceph-rest-api

VOLUME /etc/ceph

EXPOSE 9000

CMD []
ENTRYPOINT ["/gopath/bin/go-ceph-rest-api"]
