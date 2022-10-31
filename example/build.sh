#!/bin/sh

rm -Rf outer
mkdir -p outer/logs
mkdir -p outer/conf
mkdir -p outer/bin
cp conf/* outer/conf/
echo "下载扩展"
go mod tidy
echo "编译程序"
go build -o outer/bin/main main.go router.go
echo "程序开始执行"
./outer/bin/main -t=dev