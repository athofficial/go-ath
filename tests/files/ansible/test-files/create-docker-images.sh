#!/bin/bash -x

# creates the necessary docker images to run testrunner.sh locally

docker build --tag="atheios/cppjit-testrunner" docker-cppjit
docker build --tag="atheios/python-testrunner" docker-python
docker build --tag="atheios/go-testrunner" docker-go
