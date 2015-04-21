#!/bin/sh

docker build -t dumbdock-builder build
docker run -id dumbdock-builder
ID=`docker ps -ql`
rm -rf ./dumbdock
docker cp ${ID}:/go/src/app/dumbdock ./
docker rm -f $ID
