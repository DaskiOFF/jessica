#!/bin/bash

set -eux

export GOPATH="$(pwd)/.gobuild"
SRCDIR="${GOPATH}/src/github.com/daskioff/jessica"

[ -d ${GOPATH} ] && rm -rf ${GOPATH}

mkdir -p ${GOPATH}/{src,pkg,bin}
mkdir -p ${SRCDIR}
cp main.go ${SRCDIR}
cp -r ./{configs,flows,router,utils} ${SRCDIR}
(
    echo ${GOPATH}
    cd ${SRCDIR}
    go get .
    go install .
)