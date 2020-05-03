#!/bin/bash
protoc -I=./protocol -I/usr/lib -I./third_party/googleapis --go_out=plugins=grpc,paths=source_relative:. ./protocol/core/*.proto ./protocol/core/contract/*.proto 
protoc -I=./protocol -I/usr/lib -I./third_party/googleapis --go_out=plugins=grpc,paths=source_relative:. ./protocol/api/*.proto

protoc -I=./protocol -I./protocol_util -I/usr/lib -I./third_party/googleapis --go_out=plugins=grpc,paths=source_relative:./protocol_util ./protocol_util/*.proto