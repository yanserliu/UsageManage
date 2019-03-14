#!/bin/sh
rm -rf usage/
rm -rf usage.tar.gz
mkdir -p usage/
cp usage-api usage/
cp -r conf/ usage/
cp -r swagger/ usage/
tar czvf usage.tar.gz usage/
docker build -t usage-api:latest .
docker stop usage-api
docker rm usage-api
docker run -d --network=host --restart=always --name usage-api  usage-api:latest
docker ps -a | grep usage-api