#!/bin/bash

cp lib/buildinfo/buildinfo.go.in lib/buildinfo/buildinfo.go
go generate github.com/niccoloCastelli/gohls/lib/api
go run *.go ${@:1}