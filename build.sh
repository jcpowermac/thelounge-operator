#!/bin/sh

PROJECT_NAME=thelounge-operator
BIN_DIR=./build/_output/bin
BUILD_PATH="./cmd/manager"
CONTAIN_BIN_DIR=./build/bin
DLV=`which dlv`

cp $DLV ${CONTAIN_BIN_DIR}

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -gcflags='all=-N -l' -o ${BIN_DIR}/${PROJECT_NAME}-debug ${BUILD_PATH}

