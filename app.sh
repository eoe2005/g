#!/bin/sh
if [ "$1" = "" ];
then
    echo "请输入创建的路径"
    exit
fi

if [ -d $1 ];
then
    echo "文件已经存在"
    exit
else
    mkdir $1
fi

cp -r example/* $1