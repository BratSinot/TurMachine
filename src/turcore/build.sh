#!/bin/bash

if [ "$1" == "clean" ]; then
	rm -r ./build
	exit 1
fi

OLD=$PWD
cd ../../
export GOPATH=$PWD || exit 1
cd $OLD
#go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
go run turcore.go
#sudo ./build/turcore.run