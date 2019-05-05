#!/bin/bash

make_inst_dirs() {
    files=${1}/*
    for file in $files
    do
        if [ -d "$file" ]
        then
            echo install -d /app/usr/local/go"${file//.}"
            make_inst_dirs $file
        fi
    done
}

install_go() {
    files=${1}/*
    for file in $files
    do
        if [ -d "$file" ]
        then
            echo install -D $file /app/usr/local/go"${file//.}"
        fi
    done
}


install -d /app/usr/local/go
make_inst_dirs .
install_go .
ls -al /app/usr/local/go/