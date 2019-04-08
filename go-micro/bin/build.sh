#!/bin/bash

dir=`pwd`

build() {
	echo "Building $1"
	pushd $dir/$1 >/dev/null
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o "app/$1"
	popd >/dev/null
}

build api
build hash
