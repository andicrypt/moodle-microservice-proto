#!/bin/bash

DIR="golang"
RELEASE_VERSION=$1
USER_NAME=$2
EMAIL=$3
GH_TOKEN=$4
GH_REPOSITORY=$5

SERVICES=("auth" "course" "payment" "user")

git config user.name "$USER_NAME"
git config user.email "$EMAIL"
git fetch --all && git checkout main

sudo apt-get install -y protobuf-compiler golang-goprotobuf-dev
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

if [ ! -d "$DIR" ]; then
  echo "Directory $DIR does not exist. Creating it..."
  mkdir -p "$DIR"
else
  echo "Directory $DIR already exists."
fi


for SERVICE_NAME in "${SERVICES[@]}"; do
    protoc --go_out=./golang --go_opt=paths=source_relative \
    --go-grpc_out=./golang --go-grpc_opt=paths=source_relative \
    ./${SERVICE_NAME}/*.proto
    cd golang/${SERVICE_NAME}
    go mod init \
    github.com/andicrypt/microservices-proto/golang/${SERVICE_NAME} || true
    go mod tidy
    cd ../../
 
 
done

git remote set-url origin https://x-access-token:${GH_TOKEN}@github.com/${GH_REPOSITORY}
git remote -v

git add .
git commit -m "Proto update - ${RELEASE_VERSION}" || echo "No changes to commit"
git push origin HEAD:main
git tag -fa "golang/${RELEASE_VERSION}" -m "golang/${RELEASE_VERSION}"
git push origin "refs/tags/golang/${RELEASE_VERSION}"