#!/bin/bash

RUN_PATH=`pwd`
sed -i "s|RUN_PATH|$RUN_PATH|g" docker-compose.yml

#echo $RUN_PATH
docker-compose up -d --build
