#!/bin/bash

PROTOC_VERSION=26.1
PROTOC_URL="https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip"
PROTOC="./protoc/bin/protoc"
export PATH="$PATH:$HOME/go/bin"

if [ ! -f $PROTOC ]; then
  echo "Downloading protoc"
  wget -O protoc.zip $PROTOC_URL
  unzip protoc.zip -d protoc
  rm protoc.zip
fi

PROTOC="$PROTOC -I ./messages"

PROTO_SRC="./messages/*.proto  ./messages/data/*.proto"

echo "Generating protobuf files for Go"
GO_CODE_PATH=./go
if [ ! -d $GO_CODE_PATH ];then
  mkdir $GO_CODE_PATH
fi
# shellcheck disable=SC2086
$PROTOC --go_out=./go --go_opt=paths=source_relative --go-grpc_out=$GO_CODE_PATH --go-grpc_opt=paths=source_relative $PROTO_SRC
