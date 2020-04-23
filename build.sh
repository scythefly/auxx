#! /bin/bash

DIR="$( cd "$( dirname "$0" )" && pwd )"

go build -v -a -ldflags="-w -s " -trimpath .
