#!/usr/bin/env bash

protoc greet/greet.proto --go_out=plugins=grpc:.
