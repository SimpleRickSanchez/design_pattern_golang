#!/bin/bash

for dir in */; do    
    dirname=${dir%/}
    cd "$dirname" || exit
    go mod init "$dirname"
    cd - || exit
done
s
for dir in */; do
    dirname=${dir%/}
    go work use "./$dirname"
done