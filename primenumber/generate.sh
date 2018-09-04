#!/bin/bash

protoc primenumber/decomposepb/decompose.proto --go_out=plugins=grpc:.
