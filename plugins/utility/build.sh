#! /bin/bash

DIR="$(cd "$(dirname "$0")" && pwd)"
OUTPUT="$1"

if [ -z "$1" ]
then
  OUTPUT="${DIR}/utility.so"
fi

go build -buildmode=plugin -o ${OUTPUT} ${DIR}/.