#!/bin/sh

rm -Rf outer
mkdir -p outer/logs
mkdir -p outer/conf
mkdir -p outer/bin
cp conf/* outer/conf/
go mod tidy -go=1.17
go build -o outer/bin/main main.go