@echo off

DEL -r outer
mkdir  outer\logs
mkdir  outer\conf
mkdir  outer\bin
COPY  conf/* outer/conf/
echo "go mod tidy"
go mod tidy
echo "build go"
go build -o outer\bin\main.exe main.go router.go
echo "go run"

.\outer\bin\main.exe