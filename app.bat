@echo off

IF "%~1" == "" (
  echo "place input dirname"
  exit
)

IF exist "%~1" (
  echo "file exists"
  exit
) ELSE (
  mkdir "%~1"
)
XCOPY   "example/*" "%~1" /S /F