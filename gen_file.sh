#!/bin/bash

SERVICE_NAME="order"

mkdir ${SERVICE_NAME}
touch ${SERVICE_NAME}/${SERVICE_NAME}.proto

echo -e "syntax = \"proto3\";\n\noption go_package = \"github.com/andicrypt/moodle-microservice-proto/${SERVICE_NAME}\";" > ${SERVICE_NAME}/${SERVICE_NAME}.proto
